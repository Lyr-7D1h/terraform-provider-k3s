package k3s

import "io/ioutil"


type k3s struct {
	public_ssh_key string
}

func New(public_ssh_key_path string) (k3s, error) {
	public_ssh_key, err := ioutil.ReadFile(public_ssh_key_path)

	if err != nil {
		panic(err)	
	}

	k := k3s {
		public_ssh_key: string(public_ssh_key),
	}
	return k, nil
}

func (s k3s) connect() {

}