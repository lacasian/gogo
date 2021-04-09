package confgen

import (
	"log"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func Viper(c map[string]interface{}, cmd *cobra.Command, ignore []string) ([]byte, error) {
	log.Println("NOTICE: auto generated config might contain sensitive information")
	_ = cleanC(c, "", ignore)

	b, err := yaml.Marshal(c)
	if err != nil {
		return nil, errors.Wrap(err, "marshal yaml")
	}

	var node yaml.Node
	err = yaml.Unmarshal(b, &node)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal yaml")
	}

	parseNode(node.Content, "", cmd)

	b, err = yaml.Marshal(node.Content[0])
	if err != nil {
		return nil, errors.Wrap(err, "re-marshal yaml")
	}

	return b, nil
}

func parseNode(nodes []*yaml.Node, key string, cmd *cobra.Command) {
	var lastKey string
	for _, n := range nodes {
		if n.Value != "" {
			lastKey = n.Value
			if key != "" {
				lastKey = key + "." + lastKey
			}

			f := cmd.Flag(lastKey)
			if f != nil {
				n.HeadComment = f.Usage
			}
		}
		if n.Content != nil {
			parseNode(n.Content, lastKey, cmd)
		}
	}
}

func cleanC(c interface{}, key string, ignore []string) bool {
	switch tc := c.(type) {
	case map[string]interface{}:
		for k, v := range tc {
			del := cleanC(v, k, ignore)
			if del {
				delete(tc, k)
			}
		}
	default:
		for _, i := range ignore {
			if key == i {
				return true
			}
		}
	}
	return false
}
