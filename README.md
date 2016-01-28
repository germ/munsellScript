# munsellScript
MunsellPicker: Because Hex codes are not meant for humans.

MunsellPicker is a simple way to transform sRGB colours into Munsell references and back again. Turns out, RGB is terrible for humans to wrangle their head around.

The Munsell system deconstructs a colour into three components, it’s Hue, Value and Chromacity. The only difference between 5P 5/10 and 5P 6/10, is one level of Value, the other two factors are left untouched! A little useful.

Here it is in action

![MunsellPicker](https://raw.githubusercontent.com/germ/munsellScript/master/demo.gif)

MunsellPicker is simple to install.

- Place “Munsell” in /Applications (might have to right click run it once)
- Double click the service to install it
- Open System Prefs > Keyboard > Shortcuts > Text and assign munService to a shortcut

To use the picker, simply highlight the text and press your combo. If the text can be edited, it will be replaced with a corresponding Munsell value. Otherwise it will be copied to your clipboard. 

Note:
For some reason text selected Safari will not be auto copied. Just manually copy it and invoke MunsellPicker, the result will be put in your clipboard. (Although if Safari is focused, I can’t seem to invoke MunsellPicker).

Anyway, thanks for checking this grease out.
