package mustard

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

//CreateInputWidget - Creates and returns a new Input Widget
func CreateInputWidget() *InputWidget {
	var widgets []interface{}

	return &InputWidget{
		widget: widget{

			needsRepaint: true,
			widgets:      widgets,

			ref: "input",

			cursor: glfw.CreateStandardCursor(glfw.IBeamCursor),

			backgroundColor: "#fff",
		},

		fontSize:  20,
		fontColor: "#000",
	}
}

//AttachWidget - Attaches a new widget to the window
func (input *InputWidget) AttachWidget(widget interface{}) {
	input.widgets = append(input.widgets, widget)
}

//SetWidth - Sets the input width
func (input *InputWidget) SetWidth(width int) {
	input.box.width = width
	input.fixedWidth = true
}

//SetHeight - Sets the input height
func (input *InputWidget) SetHeight(height int) {
	input.box.height = height
	input.fixedHeight = true
}

//SetFontSize - Sets the input font size
func (input *InputWidget) SetFontSize(fontSize float64) {
	input.fontSize = fontSize
	input.needsRepaint = true
}

func (input *InputWidget) SetReturnCallback(returnCallback func()) {
	input.returnCallback = returnCallback
}

//SetFontColor - Sets the input font color
func (input *InputWidget) SetFontColor(fontColor string) {
	if len(fontColor) > 0 && string(fontColor[0]) == "#" {
		input.fontColor = fontColor
		input.needsRepaint = true
	}
}

//SetFontColor - Sets the input font color
func (input *InputWidget) SetValue(value string) {
	input.value = value
}

//SetFontColor - Sets the input font color
func (input *InputWidget) GetValue() string {
	return input.value
}

//SetBackgroundColor - Sets the input background color
func (input *InputWidget) SetBackgroundColor(backgroundColor string) {
	if len(backgroundColor) > 0 && string(backgroundColor[0]) == "#" {
		input.backgroundColor = backgroundColor
		input.needsRepaint = true
	}
}

func (input *InputWidget) draw(window *Window) {
	top, left, width, height := input.computedBox.GetCoords()
	context := window.context

	if input.selected {
		context.SetHexColor("#e4e4e4")
	} else {
		context.SetHexColor("#efefef")
	}

	if input.active {
		context.SetHexColor("#fff")
	}

	context.DrawRectangle(
		float64(left)+input.padding,
		float64(top)+input.padding,
		float64(width)-(input.padding*2),
		float64(height)-(input.padding*2),
	)

	context.Fill()

	context.SetHexColor("#000")
	context.SetLineWidth(.4)

	context.DrawRectangle(
		float64(left)+1+input.padding,
		float64(top)+1+input.padding,
		float64(width)-2-(input.padding*2),
		float64(height)-2-(input.padding*2),
	)

	context.SetLineJoinRound()
	context.Stroke()

	context.SetHexColor("#2f2f2f")
	context.LoadFontFace("roboto.ttf", input.fontSize)
	context.DrawString(input.value, float64(left)+input.fontSize/4+4, float64(top)+float64(height)/2+2+input.fontSize/4)
	context.Fill()

	if input.active {
		w, _ := context.MeasureString(input.value)

		context.SetHexColor("#000")
		context.DrawRectangle(
			float64(left)+input.fontSize/4+4+w,
			float64(top)+float64(height)/2-input.fontSize/2+.5,
			1.3,
			float64(input.fontSize),
		)
		context.Fill()
	}
}
