package shell

import (
  original_shell "github.com/ipfs/go-ipfs-api"
)

func RecusiveSize(s *original_shell.Shell, obj *original_shell.LsLink) (*original_shell.LsLink, error) {
  list, err := s.List(obj.Hash)
  if err != nil {
    return obj, err
  }

  if len(list) == 0 {
    return obj, nil
  }

  obj.Size = 0
  for _, sub_obj := range list {
    updated_sub_obj, err := RecusiveSize(s, sub_obj)
    if err != nil {
      return obj, err
    }

    obj.Size += updated_sub_obj.Size
  }

  return obj, nil
}

func List(s *original_shell.Shell, path string) ([]*original_shell.LsLink, error) {
  list, err := s.List(path)
  if err != nil {
    return list, err
  }

  for i, sub_obj := range list {
    list[i], err = RecusiveSize(s, sub_obj)
    if err != nil {
      return list, err
    }
  }

  return list, nil
}
