package gui

import (
	"github.com/gotk3/gotk3/gtk"
)

// NewMainWindow creates the main Window
func NewMainWindow() (*MainControl, error) {
	box, boxErr := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 5)
	if boxErr != nil {
		return nil, boxErr
	}

	streamWidget, streamWidgetErr := NewStreamControl()
	if streamWidgetErr != nil {
		return nil, streamWidgetErr
	}
	box.PackStart(streamWidget.Root(), true, true, 2)

	controlsBox, pump, camera, region, cmplxErr := NewPumpAndRegionContainer()
	if cmplxErr != nil {
		return nil, cmplxErr
	}
	box.PackEnd(controlsBox, false, true, 2)

	box.ShowAll()
	return &MainControl{rootBox: box, StreamControl: streamWidget, RegionStream: region, Camera: camera, Pump: pump}, nil
}

// MainControl is the root widget of the main window
type MainControl struct {
	Control

	// Contained elements of main window
	StreamControl *StreamControl
	RegionStream  *RegionControl
	Pump          *PumpControl
	Camera        *CameraControl

	// HBox is the hozirontal layout
	// Stream window and device selectors on the left
	// Pump controls on the right
	rootBox *gtk.Box
}

// Root returns the box that contains everything
func (m *MainControl) Root() gtk.IWidget {
	return m.rootBox
}
