package pack

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)



type ZipPacker struct {
	DirName			string
	DstName			string
}

func (z *ZipPacker) init() error {
	_, err := os.Stat(z.DirName)
	if err != nil {
		return err
	}

	if _, err := os.Stat(z.DstName); os.IsNotExist(err){
		return nil
	}

	return os.RemoveAll(z.DstName)
}

func (z *ZipPacker) Pack() error {
	if err := z.init(); err != nil{
		return err
	}

	file, err := os.Create(z.DstName)
	if err != nil {
		return err
	}
	defer file.Close()

	zw := zip.NewWriter(file)
	defer zw.Close()

	return filepath.Walk(z.DirName, func(p string, i os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		fHeader, err := zip.FileInfoHeader(i)
		if err != nil {
			return err
		}
		//fHeader.Name = strings.TrimPrefix(p, string(filepath.Separator))
		fHeader.Name = filepath.Base(p)
		if i.IsDir() {
			fHeader.Name += "/"
		} else {
			fHeader.Method = zip.Deflate
		}

		w, err := zw.CreateHeader(fHeader)
		if err != nil {
			return err
		}
		if !i.Mode().IsRegular() {
			return nil
		}
		r, err := os.Open(p)
		if err != nil {
			return err
		}
		defer r.Close()

		if _, err := io.Copy(w, r); err != nil{
			return err
		}

		return nil
	})
}

func (z *ZipPacker) UnPack() error {
	return nil
}
