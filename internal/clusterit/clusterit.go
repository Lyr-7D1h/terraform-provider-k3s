package clusterit

import "io/ioutil"


type clusterit struct {
	public_ssh_key string
}

func New(public_ssh_key_path string) (clusterit, error) {
	public_ssh_key, err := ioutil.ReadFile(public_ssh_key_path)

	if err != nil {
		panic(err)	
	}

	k := clusterit {
		public_ssh_key: string(public_ssh_key),
	}
	return k, nil
}

func (s clusterit) connect() {

}