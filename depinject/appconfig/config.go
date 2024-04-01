package appconfig

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-proto/anyutil"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"sigs.k8s.io/yaml"

	appv1alpha1 "cosmossdk.io/api/cosmos/app/v1alpha1"

	"cosmossdk.io/depinject"
	internal "cosmossdk.io/depinject/internal/appconfig"
)

// LoadJSON loads an app config in JSON format.
func LoadJSON(bz []byte) depinject.Config {
	config := &appv1alpha1.Config{}
	err := protojson.Unmarshal(bz, config)
	if err != nil {
		return depinject.Error(err)
	}

	return Compose(config)
}

// LoadYAML loads an app config in YAML format.
func LoadYAML(bz []byte) depinject.Config {
	j, err := yaml.YAMLToJSON(bz)
	if err != nil {
		return depinject.Error(err)
	}

	return LoadJSON(j)
}

// WrapAny marshals a proto message into a proto Any instance
func WrapAny(config protoreflect.ProtoMessage) *anypb.Any {
	cfg, err := anyutil.New(config)
	if err != nil {
		panic(err)
	}

	return cfg
}

// Compose composes a v1alpha1 app config into a container option by resolving
// the required modules and composing their options.
func Compose(appConfig *appv1alpha1.Config) depinject.Config {
	opts := []depinject.Config{
		depinject.Supply(appConfig),
	}

	moduleRegistry, err := internal.ModulesByModuleTypeName()
	if err != nil {
		return depinject.Error(err)
	}

	for _, module := range appConfig.Modules {
		if module.Name == "" {
			return depinject.Error(fmt.Errorf("module is missing name"))
		}

		moduleType := module.TypeUrl
		if moduleType == "" {
			if module.Config == nil {
				return depinject.Error(fmt.Errorf("module %q is missing a type and a config", module.Name))
			}

			moduleType = module.Config.TypeUrl
		}

		init, ok := moduleRegistry[moduleType]
		if !ok {
			// TODO:
			//modDesc := proto.GetExtension(msgType.Descriptor().Options(), appv1alpha1.E_Module).(*appv1alpha1.ModuleDescriptor)
			//if modDesc == nil {
			//	return depinject.Error(fmt.Errorf("no module registered for type URL %s and that protobuf type does not have the option %s\n\n%s",
			//		module.Config.TypeUrl, appv1alpha1.E_Module.TypeDescriptor().FullName(), dumpRegisteredModules(moduleRegistry)))
			//}

			return depinject.Error(fmt.Errorf("no module registered for type URL %s, did you forget to import %s: find more information on how to make a module ready for app wiring: https://docs.cosmos.network/main/building-modules/depinject\n\n%s",
				module.Config.TypeUrl, modDesc.GoImport, dumpRegisteredModules(moduleRegistry)))
		}

		if module.Config == nil {
			return depinject.Error(fmt.Errorf("module %q is missing a config object", module.Name))
		}

		config := init.ConfigProtoMessage.ProtoReflect().Type().New().Interface()
		err = anypb.UnmarshalTo(module.Config, config, proto.UnmarshalOptions{})
		if err != nil {
			return depinject.Error(err)
		}

		opts = append(opts, depinject.Supply(config))

		for _, provider := range init.Providers {
			opts = append(opts, depinject.ProvideInModule(module.Name, provider))
		}

		for _, invoker := range init.Invokers {
			opts = append(opts, depinject.InvokeInModule(module.Name, invoker))
		}

		for _, binding := range module.GolangBindings {
			opts = append(opts, depinject.BindInterfaceInModule(module.Name, binding.InterfaceType, binding.Implementation))
		}
	}

	for _, binding := range appConfig.GolangBindings {
		opts = append(opts, depinject.BindInterface(binding.InterfaceType, binding.Implementation))
	}

	return depinject.Configs(opts...)
}

func dumpRegisteredModules(modules map[string]*internal.ModuleInitializer) string {
	var mods []string
	for name := range modules {
		mods = append(mods, "  "+name)
	}
	return fmt.Sprintf("registered modules are:\n%s", strings.Join(mods, "\n"))
}
