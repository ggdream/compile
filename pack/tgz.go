package pack

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
)



type TgzPacker struct {
	DirName			string
	DstName			string
}

func (z *TgzPacker) init() error {
	_, err := os.Stat(z.DirName)
	if err != nil {
		return err
	}

	if _, err := os.Stat(z.DstName); os.IsNotExist(err){
		return nil
	}

	return os.RemoveAll(z.DstName)
}

func (z *TgzPacker) Pack() error {
	if err := z.init(); err != nil{
		return err
	}

	file, err := os.Create(z.DstName)
	if err != nil {
		return err
	}
	defer file.Close()

	gw := gzip.NewWriter(file)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()


	return filepath.Walk(z.DirName, func(p string, i os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		fHeader, err := tar.FileInfoHeader(i, "conf")
		if err != nil {
			return err
		}
		fHeader.Name = filepath.Base(p)

		if err := tw.WriteHeader(fHeader); err != nil {
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

		if _, err := io.Copy(tw, r); err != nil{
			return err
		}

		return nil
	})
}

func (z *TgzPacker) UnPack() error {
	return nil
}
