	"github.com/m3db/m3/src/dbnode/encoding/proto"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
	var schema *desc.MessageDescriptor
	if cfg.DataMode == storage.DataModeProtoBuf {
		logger.Info("Probuf data mode enabled")
		schema, err = parseProtoSchema(cfg.Proto.SchemaFilePath)
		if err != nil {
			logger.Fatalf("error parsing protobuffer schema: %v", err)
		}
	}

		},
		func(opts client.AdminOptions) client.AdminOptions {
			if cfg.DataMode == storage.DataModeProtoBuf {
				return opts.SetEncodingProto(schema, encoding.NewOptions())
			}
			return opts
		},
	)
	schema *desc.MessageDescriptor,
		if schema != nil {
			// TODO: should probably allow a schema to be passed on construction.
			enc := proto.NewEncoder(time.Time{}, encodingOpts)
			enc.SetSchema(schema)
			return enc
		}

		return m3tsz.NewEncoder(time.Time{}, nil, nil, m3tsz.DefaultIntOptimizationEnabled, encodingOpts)
		if schema != nil {
			return proto.NewIterator(r, schema, encodingOpts)
		}
		return m3tsz.NewReaderIterator(r, nil, m3tsz.DefaultIntOptimizationEnabled, encodingOpts)
		SetReaderIteratorPool(iteratorPool).

func parseProtoSchema(filePath string) (*desc.MessageDescriptor, error) {
	fds, err := protoparse.Parser{}.ParseFiles(filePath)
	if err != nil {
		return nil, fmt.Errorf(
			"error parsing proto schema: %s, err: %v", filePath, err)
	}

	if len(fds) != 1 {
		return nil, fmt.Errorf(
			"expected to parse %s into one file descriptor but parse: %s",
			filePath, len(fds))
	}

	// TODO(rartoul): This will be more sophisticated later, but for now assume
	// that the message will be called "Schema".
	schema := fds[0].FindMessage("Schema")
	if schema == nil {
		return nil, fmt.Errorf(
			"expected to find message with name 'Schema' in %s, but did not",
			filePath,
		)
	}

	return schema, nil
}