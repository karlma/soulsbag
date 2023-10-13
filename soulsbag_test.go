package soulsbag

import (
	"sync"
	"testing"

	"github.com/pmezard/go-difflib/difflib"

	"github.com/karlma/soulsbag/source"
)

type (
	MySQLCnf struct {
		Host         string `toml:"host"`
		Port         int
		User         string
		Password     string
		DBName       string `toml:"db_name" yaml:"db_name" json:"db_name"`
		MaxIdleConns int    `toml:"max_idle_conns" yaml:"max_idle_conns" json:"max_idle_conns"`
		MaxOpenConns int    `toml:"max_open_conns" yaml:"max_open_conns" json:"max_open_conns"`
		EnableLog    bool   `toml:"enable_log" yaml:"enable_log" json:"enable_log"`
	}
	TowerNewServerCnf struct {
		Domain string   `toml:"domain" yaml:"domain"`
		Apps   []string `toml:"apps" yaml:"apps"`
	}

	TowerCnf struct {
		ListenIp   string            `toml:"listen_ip" yaml:"listen_ip" json:"listen_ip"`
		ListenPort int               `toml:"listen_port" yaml:"listen_port" json:"listen_port"`
		NewServer  TowerNewServerCnf `toml:"new_server" yaml:"new_server" json:"new_server"`
	}

	MyCnf struct {
		MySQL MySQLCnf `toml:"mysql" yaml:"mysql"`
		Tower TowerCnf `toml:"tower" yaml:"tower"`
	}
)

func TestSoulsbagFile(t *testing.T) {
	t.Run("toml", func(t *testing.T) {
		fileT(t, "./testdata.toml", "toml")
	})
	t.Run("toml", func(t *testing.T) {
		fileT_2(t, "./testdata.toml", "toml")
	})
	t.Run("yaml", func(t *testing.T) {
		fileT(t, "./testdata.yaml", "yaml")
	})
	t.Run("yaml", func(t *testing.T) {
		fileT_2(t, "./testdata.yaml", "yaml")
	})
	t.Run("json", func(t *testing.T) {
		fileT(t, "./testdata.json", "json")
	})
	t.Run("json", func(t *testing.T) {
		fileT_2(t, "./testdata.json", "json")
	})
}

func fileT(t *testing.T, fName, fTyp string) {
	err := Init("file", fTyp, source.Options{
		Path: fName,
	})
	if err != nil {
		t.Fatalf("init config error: %v", err)
	}

	err = Read()
	if err != nil {
		t.Fatalf("read config error: %v", err)
	}

	var cfg MyCnf
	err = Unmarshal(&cfg)
	if err != nil {
		t.Fatalf("unmarshal config error: %v", err)
	}

	//t.Logf("unmarshal: %#v\n", cfg)
	if cfg.MySQL.Host != "192.168.1.112" {
		t.Errorf("read mysqlCfg.Host: %s, expected: %s", cfg.MySQL.Host, "192.168.1.112")
	}

	if cfg.MySQL.EnableLog != false {
		t.Errorf("read mysqlCfg.EnableLog: %t, expected: %t", cfg.MySQL.EnableLog, false)
	}

	if cfg.Tower.ListenPort != 4022 {
		t.Errorf("read towerCfg.ListenPort: %d, expected: %d", cfg.Tower.ListenPort, 4022)
	}

	if cfg.Tower.NewServer.Domain != "xxx.example.com" {
		t.Errorf("read towerCfg.NewServer.Domain: %s, expected: %s",
			cfg.Tower.NewServer.Domain, "xxx.example.com")
	}
}

func fileT_2(t *testing.T, fName, fTyp string) {
	sb := New()
	err := sb.Init("file", fTyp, source.Options{
		Path: fName,
	})
	if err != nil {
		t.Fatalf("init config error: %v", err)
	}

	err = sb.Read()
	if err != nil {
		t.Fatalf("read config error: %v", err)
	}

	var cfg MyCnf
	err = sb.Unmarshal(&cfg)
	if err != nil {
		t.Fatalf("unmarshal config error: %v", err)
	}

	//t.Logf("unmarshal: %#v\n", cfg)
	if cfg.MySQL.Host != "192.168.1.112" {
		t.Errorf("read mysqlCfg.Host: %s, expected: %s", cfg.MySQL.Host, "192.168.1.112")
	}

	if cfg.MySQL.EnableLog != false {
		t.Errorf("read mysqlCfg.EnableLog: %t, expected: %t", cfg.MySQL.EnableLog, false)
	}

	if cfg.Tower.ListenPort != 4022 {
		t.Errorf("read towerCfg.ListenPort: %d, expected: %d", cfg.Tower.ListenPort, 4022)
	}

	if cfg.Tower.NewServer.Domain != "xxx.example.com" {
		t.Errorf("read towerCfg.NewServer.Domain: %s, expected: %s",
			cfg.Tower.NewServer.Domain, "xxx.example.com")
	}
}

func TestSoulsBagEtcdv3(t *testing.T) {
	t.Run("toml", func(t *testing.T) {
		etcdv3T(t, "toml")
	})
	t.Run("toml", func(t *testing.T) {
		etcdv3T_2(t, "toml")
	})

	t.Run("yaml", func(t *testing.T) {
		etcdv3T(t, "yaml")
	})
	t.Run("yaml", func(t *testing.T) {
		etcdv3T_2(t, "yaml")
	})
}

func etcdv3T(t *testing.T, encTyp string) {
	err := Init("etcdv3", encTyp, source.Options{
		Path:         "127.0.0.1:2379",
		Key:          "cs/s4/soulsbag." + encTyp,
		AuthUser:     "root",
		AuthPassword: "123456",
	})
	if err != nil {
		t.Errorf("init config error: %v", err)
	}

	err = Read()
	if err != nil {
		t.Fatalf("read config error: %v", err)
	}

	var cfg MyCnf
	err = Unmarshal(&cfg)
	if err != nil {
		t.Fatalf("unmarshal config error: %v", err)
	}

	//t.Logf("unmarshal: %#v\n", cfg)
	if cfg.MySQL.Host != "192.168.1.112" {
		t.Errorf("read mysqlCfg.Host: %s, expected: %s", cfg.MySQL.Host, "192.168.1.112")
	}

	if cfg.MySQL.EnableLog != false {
		t.Errorf("read mysqlCfg.EnableLog: %t, expected: %t", cfg.MySQL.EnableLog, false)
	}

	if cfg.Tower.ListenPort != 4022 {
		t.Errorf("read towerCfg.ListenPort: %d, expected: %d", cfg.Tower.ListenPort, 4022)
	}

	if cfg.Tower.NewServer.Domain != "xxx.example.com" {
		t.Errorf("read towerCfg.NewServer.Domain: %s, expected: %s",
			cfg.Tower.NewServer.Domain, "xxx.example.com")
	}
}

func etcdv3T_2(t *testing.T, encTyp string) {
	sb := New()

	err := sb.Init("etcdv3", encTyp, source.Options{
		Path:         "127.0.0.1:2379",
		Key:          "cs/s4/soulsbag." + encTyp,
		AuthUser:     "root",
		AuthPassword: "123456",
	})
	if err != nil {
		t.Errorf("init config error: %v", err)
	}

	err = sb.Read()
	if err != nil {
		t.Fatalf("read config error: %v", err)
	}

	var cfg MyCnf
	err = sb.Unmarshal(&cfg)
	if err != nil {
		t.Fatalf("unmarshal config error: %v", err)
	}

	//t.Logf("unmarshal: %#v\n", cfg)
	if cfg.MySQL.Host != "192.168.1.112" {
		t.Errorf("read mysqlCfg.Host: %s, expected: %s", cfg.MySQL.Host, "192.168.1.112")
	}

	if cfg.MySQL.EnableLog != false {
		t.Errorf("read mysqlCfg.EnableLog: %t, expected: %t", cfg.MySQL.EnableLog, false)
	}

	if cfg.Tower.ListenPort != 4022 {
		t.Errorf("read towerCfg.ListenPort: %d, expected: %d", cfg.Tower.ListenPort, 4022)
	}

	if cfg.Tower.NewServer.Domain != "xxx.example.com" {
		t.Errorf("read towerCfg.NewServer.Domain: %s, expected: %s",
			cfg.Tower.NewServer.Domain, "xxx.example.com")
	}
}

func TestSoulsBagWatch(t *testing.T) {
	t.Run("toml etcdv1", func(t *testing.T) {
		encTyp := "toml"
		err := Init("etcdv3", encTyp, source.Options{
			Path:         "127.0.0.1:2379",
			Key:          "cs/s4/soulsbag." + encTyp,
			AuthUser:     "root",
			AuthPassword: "123456",
		})
		if err != nil {
			t.Errorf("init config error: %v", err)
		}

		err = Read()
		if err != nil {
			t.Fatalf("read config error: %v", err)
		}
		originalData := string(GetData())

		watchWG := sync.WaitGroup{}
		watchWG.Add(1)
		t.Logf("waiting you to change config file ...")
		err = Watch(func(event string) {
			t.Logf("watch event: [%s]", event)
			watchWG.Done()
		})
		if err != nil {
			t.Fatalf("watch config error: %v", err)
		}
		watchWG.Wait()

		// After modify, read again
		err = Read()
		if err != nil {
			t.Fatalf("read config error: %v", err)
		}

		currentData := string(GetData())

		diff := difflib.ContextDiff{
			A:        difflib.SplitLines(originalData),
			B:        difflib.SplitLines(currentData),
			FromFile: "original",
			ToFile:   "current",
			Eol:      "\n",
			Context:  2,
		}
		result, _ := difflib.GetContextDiffString(diff)
		t.Logf("diff:\n%s", result)
	})
}
