hz new -I idl -idl idl/douyin/feed.proto --out_dir services/feed --module nico_minidouyin/services/feed
#PS C:\Users\nicognaw\GolandProjects\nico_minidouyin> hz new --help
 #NAME:
 #   hz new - Generate a new Hertz project
 #
 #USAGE:
 #   hz new [command options] [arguments...]
 #
 #OPTIONS:
 #   --idl value [ --idl value ]                                            Specify the IDL file path. (.thrift or .proto)
 #   --service value                                                        Specify the service name.
 #   --module value, --mod value                                            Specify the Go module name to generate go.mod.
 #   --out_dir value                                                        Specify the project path.
 #   --handler_dir value                                                    Specify the handler path.
 #   --model_dir value                                                      Specify the model path.
 #   --client_dir value                                                     Specify the client path. If not specified, IDL generated path is used for 'client' command; no client code is generated for 'new' command
 #   --proto_path value, -I value [ --proto_path value, -I value ]          Add an IDL search path for includes. (Valid only if idl is protobuf)
 #   --thriftgo value, -t value [ --thriftgo value, -t value ]              Specify arguments for the thriftgo. ({flag}={value})
 #   --protoc value, -p value [ --protoc value, -p value ]                  Specify arguments for the protoc. ({flag}={value})
 #   --option_package value, -P value [ --option_package value, -P value ]  Specify the package path. ({include_path}={import_path})
 #   --no_recurse                                                           Generate master model only. (default: false)
 #   --json_enumstr                                                         Use string instead of num for json enums when idl is thrift. (default: false)
 #   --unset_omitempty                                                      Remove 'omitempty' tag for generated struct. (default: false)
 #   --pb_camel_json_tag                                                    Convert Name style for json tag to camel(Only works protobuf). (default: false)
 #   --snake_tag                                                            Use snake_case style naming for tags. (Only works for 'form', 'query', 'json') (default: false)
 #   --exclude_file value, -E value [ --exclude_file value, -E value ]      Specify the files that do not need to be updated.
 #   --customize_layout value                                               Specify the layout template. ({{Template Profile}}:{{Rendering Data}})
 #   --customize_package value                                              Specify the package template. ({{Template Profile}}:)
 #   --handler_by_method                                                    Generate a separate handler file for each method. (default: false)
 #   --protoc-plugins value [ --protoc-plugins value ]                      Specify plugins for the protoc. ({plugin_name}:{options}:{out_dir})
 #   --thrift-plugins value [ --thrift-plugins value ]                      Specify plugins for the thriftgo. ({plugin_name}:{options})
 #   --help, -h                                                             show help (default: false)