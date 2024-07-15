package config

import (
	"fmt"
	"strings"
	
	"github.com/iancoleman/strcase"
	
	"github.com/daarlabs/hirokit/devtool"
	. "github.com/daarlabs/hirokit/gox"
	"github.com/daarlabs/hirokit/tempest"
)

const (
	StarterRepo = "https://github.com/daarlabs/starter.git"
)

func init() {
	devtool.ToolConfig.Plugin[devtool.PluginDatabase] = devtool.Plugin{
		Title:    "Database",
		IconPath: Path(D("M5 12.5C5 12.8134 5.46101 13.3584 6.53047 13.8931C7.91405 14.5849 9.87677 15 12 15C14.1232 15 16.0859 14.5849 17.4695 13.8931C18.539 13.3584 19 12.8134 19 12.5V10.3287C17.35 11.3482 14.8273 12 12 12C9.17273 12 6.64996 11.3482 5 10.3287V12.5ZM19 15.3287C17.35 16.3482 14.8273 17 12 17C9.17273 17 6.64996 16.3482 5 15.3287V17.5C5 17.8134 5.46101 18.3584 6.53047 18.8931C7.91405 19.5849 9.87677 20 12 20C14.1232 20 16.0859 19.5849 17.4695 18.8931C18.539 18.3584 19 17.8134 19 17.5V15.3287ZM3 17.5V7.5C3 5.01472 7.02944 3 12 3C16.9706 3 21 5.01472 21 7.5V17.5C21 19.9853 16.9706 22 12 22C7.02944 22 3 19.9853 3 17.5ZM12 10C14.1232 10 16.0859 9.58492 17.4695 8.89313C18.539 8.3584 19 7.81342 19 7.5C19 7.18658 18.539 6.6416 17.4695 6.10687C16.0859 5.41508 14.1232 5 12 5C9.87677 5 7.91405 5.41508 6.53047 6.10687C5.46101 6.6416 5 7.18658 5 7.5C5 7.81342 5.46101 8.3584 6.53047 8.89313C7.91405 9.58492 9.87677 10 12 10Z")),
		RowFunc: func(value string) Node {
			parts := strings.Split(value, "/")
			return Div(
				Span(tempest.Class().Name(devtool.DynamicStyle).FontBold().TextBlue(500), Text(fmt.Sprintf("[%s] ", parts[0]))),
				Span(tempest.Class().Name(devtool.DynamicStyle).TextSlate(300), Raw(devtool.FormatSql(parts[1]))),
			)
		},
	}
	devtool.ToolConfig.Plugin[devtool.PluginDebug] = devtool.Plugin{
		Title:    "Debug",
		IconPath: Path(D("M10.5621 4.14785C11.0262 4.05095 11.5071 4.00001 12 4.00001C12.4929 4.00001 12.9738 4.05095 13.4379 4.14785L15.1213 2.46448L16.5355 3.87869L15.4859 4.92834C16.7177 5.6371 17.7135 6.70996 18.3264 8.00001H21V10H18.9291C18.9758 10.3266 19 10.6605 19 11V12H21V14H19V15C19 15.3395 18.9758 15.6734 18.9291 16H21V18H18.3264C17.2029 20.365 14.7924 22 12 22C9.2076 22 6.7971 20.365 5.67363 18H3V16H5.07089C5.02417 15.6734 5 15.3395 5 15V14H3V12H5V11C5 10.6605 5.02417 10.3266 5.07089 10H3V8.00001H5.67363C6.28647 6.70996 7.28227 5.6371 8.51412 4.92834L7.46447 3.87869L8.87868 2.46448L10.5621 4.14785ZM12 6.00001C9.23858 6.00001 7 8.23859 7 11V15C7 17.7614 9.23858 20 12 20C14.7614 20 17 17.7614 17 15V11C17 8.23859 14.7614 6.00001 12 6.00001ZM9 14H15V16H9V14ZM9 10H15V12H9V10Z")),
		RowFunc: func(value string) Node {
			return Text(value)
		},
	}
	devtool.ToolConfig.Plugin[devtool.PluginSession] = devtool.Plugin{
		Title:    "Session",
		IconPath: Path(D("M12 2C17.5228 2 22 6.47715 22 12C22 17.5228 17.5228 22 12 22C6.47715 22 2 17.5228 2 12C2 6.47715 6.47715 2 12 2ZM12.1597 16C10.1243 16 8.29182 16.8687 7.01276 18.2556C8.38039 19.3474 10.114 20 12 20C13.9695 20 15.7727 19.2883 17.1666 18.1081C15.8956 16.8074 14.1219 16 12.1597 16ZM12 4C7.58172 4 4 7.58172 4 12C4 13.8106 4.6015 15.4807 5.61557 16.8214C7.25639 15.0841 9.58144 14 12.1597 14C14.6441 14 16.8933 15.0066 18.5218 16.6342C19.4526 15.3267 20 13.7273 20 12C20 7.58172 16.4183 4 12 4ZM12 5C14.2091 5 16 6.79086 16 9C16 11.2091 14.2091 13 12 13C9.79086 13 8 11.2091 8 9C8 6.79086 9.79086 5 12 5ZM12 7C10.8954 7 10 7.89543 10 9C10 10.1046 10.8954 11 12 11C13.1046 11 14 10.1046 14 9C14 7.89543 13.1046 7 12 7Z")),
		RowFunc: func(value string) Node {
			parts := strings.Split(value, "/")
			return Div(
				Span(tempest.Class().Name(devtool.DynamicStyle).FontBold(), Text(parts[0]+": ")),
				Span(Text(parts[1])),
			)
		},
	}
	devtool.ToolConfig.Plugin[devtool.PluginParam] = devtool.Plugin{
		Title:     "Parameters",
		IconPath:  Path(D("M9 3V5H6V19H9V21H4V3H9ZM15 3H20V21H15V19H18V5H15V3Z")),
		Reference: true,
		RowFunc: func(value string) Node {
			parts := strings.Split(value, "/")
			return Div(
				Span(tempest.Class().Name(devtool.DynamicStyle).FontBold(), Text(strcase.ToKebab(parts[0])+": ")),
				Span(Text(parts[1])),
			)
		},
	}
}
