package main

import (
        "time"

        "gobot.io/x/gobot"
        "gobot.io/x/gobot/drivers/spi"
        "gobot.io/x/gobot/platforms/raspi"
)

func main() {
        r := raspi.NewAdaptor()
        d := spi.NewMCP3008(r)
        work := func() {
                gobot.Every(1*time.Second, func() {
                    output := d.Read(0)
                    print(output)
                })
        }

        robot := gobot.NewRobot("pi-plant",
                []gobot.Connection{r},
                []gobot.Device{led},
                work,
        )

        robot.Start()
}
