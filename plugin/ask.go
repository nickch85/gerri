package plugin

/*
usage: !ask are you sure?
usage: !ask this or that or neither
*/

import (
	"github.com/dysfn/gerri/cmd"
	"github.com/dysfn/gerri/data"
	"math/rand"
	"strings"
	"time"
)

func ReplyAsk(pm data.Privmsg, config *data.Config) (string, error) {
	msg := strings.Join(pm.Message[1:], " ")
	if strings.TrimSpace(msg) != "" {
		rand.Seed(time.Now().UnixNano())
		ors := strings.Split(msg, " or ")
		if len(ors) > 1 {
			return cmd.Privmsg(pm.Target, ors[rand.Intn(len(ors))]), nil
		}
		return cmd.Privmsg(pm.Target, [2]string{"yes!", "no..."}[rand.Intn(2)]), nil
	}
	return "", nil
}
