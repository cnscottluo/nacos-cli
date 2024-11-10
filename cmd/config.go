package cmd

import (
	"errors"
	"fmt"
	"github.com/cnscottluo/nacos-cli/internal"
	"github.com/cnscottluo/nacos-cli/internal/nacos"
	"github.com/gdamore/tcell/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

// save config to current directory
var save bool

// config file path
var configPath string

// config type
var configType string

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "config management",
	Long:  `config management.`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var configListCmd = &cobra.Command{
	Use:   "list <namespaceId>",
	Short: "list config in namespace",
	Long:  `list config in namespace.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("namespaceId is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		namespaceId := args[0]
		result, err := nacosClient.GetConfigs(namespaceId)
		internal.CheckErr(err)
		internal.TableShow([]string{"DataId", "Group", "Type"}, internal.GenData(result, func(resp nacos.ConfigResp) []string {
			return []string{
				resp.DataId,
				resp.Group,
				resp.Type,
			}
		}))
	},
}

var configGetCmd = &cobra.Command{
	Use:   "get <dataId>",
	Short: "get config",
	Long:  `get config.`,
	Args: func(cmd *cobra.Command, args []string) error {

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		dataId := args[0]
		result, err := nacosClient.GetConfig(dataId)
		internal.CheckErr(err)
		internal.ConfigShow(dataId, result)
		if save {
			internal.SaveConfig(dataId, result)
		}
	},
}

var configDeleteCmd = &cobra.Command{
	Use:   "delete <dataId>",
	Short: "delete config",
	Long:  `delete config.`,
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		dataId := args[0]
		_, err := nacosClient.DeleteConfig(dataId)
		internal.CheckErr(err)
		internal.Info("delete config %s success", dataId)
	},
}

var configPublishCmd = &cobra.Command{
	Use:   "publish <dataId>",
	Short: "publish config",
	Long:  `publish config.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("dataId is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		dataId := args[0]
		content, err := internal.ReadFile(configPath)
		internal.CheckErr(err)
		_, err = nacosClient.PublishConfig(dataId, content, configType)
		internal.CheckErr(err)
		internal.Info("publish config %s success", dataId)
	},
}

var configEditCmd = &cobra.Command{
	Use:   "edit <dataId>",
	Short: "edit config",
	Long:  `edit config.`,
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
		boxStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorPurple)

		// Initialize screen
		s, err := tcell.NewScreen()
		if err != nil {
			log.Fatalf("%+v", err)
		}
		if err := s.Init(); err != nil {
			log.Fatalf("%+v", err)
		}
		s.SetStyle(defStyle)
		s.EnableMouse()
		s.EnablePaste()
		s.Clear()

		// Draw initial boxes
		drawBox(s, 1, 1, 42, 7, boxStyle, "Click and drag to draw a box")
		drawBox(s, 5, 9, 32, 14, boxStyle, "Press C to reset")

		quit := func() {
			// You have to catch panics in a defer, clean up, and
			// re-raise them - otherwise your application can
			// die without leaving any diagnostic trace.
			maybePanic := recover()
			s.Fini()
			if maybePanic != nil {
				panic(maybePanic)
			}
		}
		defer quit()

		// Here's how to get the screen size when you need it.
		// xmax, ymax := s.Size()

		// Here's an example of how to inject a keystroke where it will
		// be picked up by the next PollEvent call.  Note that the
		// queue is LIFO, it has a limited length, and PostEvent() can
		// return an error.
		// s.PostEvent(tcell.NewEventKey(tcell.KeyRune, rune('a'), 0))

		// Event loop
		ox, oy := -1, -1
		for {
			// Update screen
			s.Show()

			// Poll event
			ev := s.PollEvent()

			// Process event
			switch ev := ev.(type) {
			case *tcell.EventResize:
				s.Sync()
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
					return
				} else if ev.Key() == tcell.KeyCtrlL {
					s.Sync()
				} else if ev.Rune() == 'C' || ev.Rune() == 'c' {
					s.Clear()
				}
			case *tcell.EventMouse:
				x, y := ev.Position()

				switch ev.Buttons() {
				case tcell.Button1, tcell.Button2:
					if ox < 0 {
						ox, oy = x, y // record location when click started
					}

				case tcell.ButtonNone:
					if ox >= 0 {
						label := fmt.Sprintf("%d,%d to %d,%d", ox, oy, x, y)
						drawBox(s, ox, oy, x, y, boxStyle, label)
						ox, oy = -1, -1
					}
				}
			}
		}
	},
}

// init publish command
func initPublish() {
	configPublishCmd.Flags().StringVarP(&configPath, "path", "p", "", "config file path")
	_ = configPublishCmd.MarkFlagRequired("path")
	configPublishCmd.Flags().StringVarP(&configType, "type", "t", "yaml", "config type (text,json,xml,yaml,html,properties)")
	configCmd.AddCommand(configPublishCmd)
}

func initGet() {
	configGetCmd.Flags().BoolVarP(&save, "save", "s", false, "save config to current directory")
	configCmd.AddCommand(configGetCmd)
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.PersistentFlags().StringP("namespaceId", "n", "", "namespaceId")
	configCmd.PersistentFlags().StringP("group", "g", "DEFAULT_GROUP", "group")
	_ = viper.BindPFlag("nacos.namespace", configCmd.Flags().Lookup("namespaceId"))
	_ = viper.BindPFlag("nacos.group", configCmd.Flags().Lookup("group"))

	configCmd.AddCommand(configListCmd)
	configCmd.AddCommand(configDeleteCmd)
	configCmd.AddCommand(configEditCmd)
	initGet()
	initPublish()
}

func drawText(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	row := y1
	col := x1
	for _, r := range []rune(text) {
		s.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}

func drawBox(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	// Fill background
	for row := y1; row <= y2; row++ {
		for col := x1; col <= x2; col++ {
			s.SetContent(col, row, ' ', nil, style)
		}
	}

	// Draw borders
	for col := x1; col <= x2; col++ {
		s.SetContent(col, y1, tcell.RuneHLine, nil, style)
		s.SetContent(col, y2, tcell.RuneHLine, nil, style)
	}
	for row := y1 + 1; row < y2; row++ {
		s.SetContent(x1, row, tcell.RuneVLine, nil, style)
		s.SetContent(x2, row, tcell.RuneVLine, nil, style)
	}

	// Only draw corners if necessary
	if y1 != y2 && x1 != x2 {
		s.SetContent(x1, y1, tcell.RuneULCorner, nil, style)
		s.SetContent(x2, y1, tcell.RuneURCorner, nil, style)
		s.SetContent(x1, y2, tcell.RuneLLCorner, nil, style)
		s.SetContent(x2, y2, tcell.RuneLRCorner, nil, style)
	}

	drawText(s, x1+1, y1+1, x2-1, y2-1, style, text)
}
