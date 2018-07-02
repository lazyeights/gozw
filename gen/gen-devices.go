package gen

import (
	"io"
)

func (g *Generator) GenerateDevices(w io.Writer) (err error) {
	_, err = writeStrings(w,
		"// THIS FILE IS AUTO-GENERATED BY ZWGEN\n",
		"// DO NOT MODIFY\n\n",
		"package cc\n\n",
		"const (\n",
	)
	if err != nil {
		return
	}
	for _, bd := range g.zwClasses.BasicDevices {
		io.WriteString(w, "\t")
		io.WriteString(w, toGoName(bd.Name))
		io.WriteString(w, " BasicDeviceType = ")
		io.WriteString(w, bd.Key)
		_, err = io.WriteString(w, "\n")
		if err != nil {
			return err
		}
	}
	_, err = io.WriteString(w, ")\n\nvar BasicDeviceTypeNames = map[BasicDeviceType]string{\n")
	if err != nil {
		return
	}
	for _, bd := range g.zwClasses.BasicDevices {
		io.WriteString(w, "\t")
		io.WriteString(w, toGoName(bd.Name))
		io.WriteString(w, ": \"")
		io.WriteString(w, bd.Help)
		_, err = io.WriteString(w, "\",\n")
		if err != nil {
			return
		}
	}
	_, err = io.WriteString(w, "}\n\nconst (\n")
	if err != nil {
		return
	}
	for _, gd := range g.zwClasses.GenericDevices {
		io.WriteString(w, "\t")
		io.WriteString(w, toGoName(gd.Name))
		io.WriteString(w, " GenericDeviceType = ")
		io.WriteString(w, gd.Key)
		_, err = io.WriteString(w, "\n")
		if err != nil {
			return err
		}
	}
	_, err = io.WriteString(w, ")\n\nvar GenericDeviceTypeNames = map[GenericDeviceType]string{\n")
	if err != nil {
		return
	}
	for _, gd := range g.zwClasses.GenericDevices {
		io.WriteString(w, "\t")
		io.WriteString(w, toGoName(gd.Name))
		io.WriteString(w, ": \"")
		io.WriteString(w, gd.Help)
		_, err = io.WriteString(w, "\",\n")
		if err != nil {
			return
		}
	}
	_, err = io.WriteString(w, "}\n\nconst (\nSpecificTypeNotUsed SpecificDeviceType = 0x00\n")
	if err != nil {
		return
	}
	for _, gd := range g.zwClasses.GenericDevices {
		for _, sd := range gd.SpecificDevices {
			if sd.Key != "0x00" {
				io.WriteString(w, "\t")
				io.WriteString(w, toGoName(sd.Name))
				io.WriteString(w, " SpecificDeviceType = ")
				io.WriteString(w, sd.Key)
				_, err = io.WriteString(w, "\n")
				if err != nil {
					return
				}
			}
		}
	}
	_, err = io.WriteString(w,
		")\n\nvar SpecificDeviceTypeNames = map[GenericDeviceType]map[SpecificDeviceType]string{\n",
	)
	if err != nil {
		return
	}
	for _, gd := range g.zwClasses.GenericDevices {
		io.WriteString(w, "\t")
		io.WriteString(w, toGoName(gd.Name))
		_, err = io.WriteString(w, ": map[SpecificDeviceType]string{\n")
		if err != nil {
			return
		}
		for _, sd := range gd.SpecificDevices {
			io.WriteString(w, "\t\t")
			io.WriteString(w, toGoName(sd.Name))
			io.WriteString(w, ": \"")
			io.WriteString(w, sd.Help)
			_, err = io.WriteString(w, "\",\n")
			if err != nil {
				return
			}
		}
		_, err = io.WriteString(w, "\t}\n")
		if err != nil {
			return
		}
	}
	_, err = io.WriteString(w, "}\n")
	return
}