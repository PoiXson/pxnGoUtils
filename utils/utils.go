package utils;

import(
	"os"
	"path"
	"errors"
);



func IsFile(file string) bool {
	info, err := os.Stat(file);
	if err != nil {
		if errors.Is(err, os.ErrNotExist) { return false; }
		panic(err);
	}
	return info.Mode().IsRegular();
}

func FindFile(file string, paths ...string) string {
	for i := range paths {
		p := path.Join(paths[i], file);
		if IsFile(p) { return p; }
	}
	return "";
}
