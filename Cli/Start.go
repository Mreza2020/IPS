package Cli

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/Mreza2020/Image_Processing_Service/Build"
	"github.com/Mreza2020/Image_Processing_Service/DB"
	Login "github.com/Mreza2020/Image_Processing_Service/login"
)

var Authentication *DB.User

// Scanner reads a single line from standard input and returns it as a string.
//
// If a non-empty format string is provided, the input is printed using
// fmt.Printf with the input value. If the format string is empty, no output
// is printed.
//
// Scanner is used throughout the CLI to collect user input interactively.
func Scanner(s string) string {
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	output := scanner.Text()
	if s == "" {
		return output
	}
	fmt.Printf(s, output)
	return output
}

// StartCli starts the interactive command-line interface for the application.
//
// It displays the welcome message, parses the initial CLI command, and then
// continuously executes commands until the application is terminated.
//
// After each command, the user is prompted to enter the next command.
func StartCli() {
	fmt.Println("!!! Welcome to the cli program !!!")

	command := flag.String("cli", "run app", "run command")
	commandD := flag.String("SerializeMode", "txt", "Save Mode")

	flag.Parse()

	switch *commandD {
	case "":
		DB.SerializeMode = DB.SerializeMode1
		fmt.Println("txt")
	case "txt":
		DB.SerializeMode = DB.SerializeMode1
		fmt.Println("txt")
	case "json":
		DB.SerializeMode = DB.SerializeMode2
		fmt.Println("json")
	}

	DB.LoadUsers()

	for {

		RunCommand(*command)
		fmt.Println("please enter another command")
		*command = Scanner("")

	}

}

// login authenticates a user using credentials entered through the CLI.
//
// The function prompts the user for a username and password, attempts to
// authenticate the credentials, and stores the authenticated user in
// Login.Authentication when authentication succeeds.
//
// If authentication fails, the function displays an error message and returns.
func login() {
	fmt.Println("Username: ")
	username := Scanner("")

	fmt.Println("Password: ")
	password := Scanner("")

	result, user := Login.Login(username, password)

	if result == "" {
		fmt.Println("Login failed.")
		return
	}
	Authentication = user

	fmt.Println("!!! Login successfully !!!")
	imagePath()
}

var ImagePath string

// imagePath prompts the user to enter the path of the image that will be used
// by image-processing commands.
//
// The provided path is stored in the package-level ImagePath variable and is
// subsequently used by commands such as compress, crop, filter, flip, format,
// resize, rotate, and watermark.
func imagePath() {
	fmt.Println("Please enter the file path")
	filePath := Scanner("")
	ImagePath = filePath
	fmt.Println("")
}

// printHelp displays the available CLI commands and their usage information.
//
// The help output describes authentication requirements, image-processing
// commands, supported options, and common examples.
func printHelp() {
	fmt.Println("Usage: cli <command> [options]")
	fmt.Println()
	fmt.Println("Image Processing CLI")
	fmt.Println("====================")
	fmt.Println()
	fmt.Println("Available commands:")
	fmt.Println()
	fmt.Println("  compress   Compress an image and reduce its file size")
	fmt.Println("  crop       Crop an image to the specified dimensions and position")
	fmt.Println("  filter     Apply a filter to an image")
	fmt.Println("  flip       Flip an image horizontally or vertically")
	fmt.Println("  format     Convert an image to another format")
	fmt.Println("  resize     Resize an image to the specified width and height")
	fmt.Println("  rotate     Rotate an image by a specified angle")
	fmt.Println("  watermark  Add an image watermark with opacity and display mode")
	fmt.Println("  path       Show and change the current image path")
	fmt.Println("  list       Show images saved for the current user")
	fmt.Println("  sign       Sign in / register a user")
	fmt.Println("  exit       Exit the application")
	fmt.Println("  help       Show help information")
	fmt.Println()

	fmt.Println("Options:")
	fmt.Println()
	fmt.Println("  -SerializeMode")
	fmt.Println("      Select the mode used to save/serialize data.")
	fmt.Println("      Supported modes:")
	fmt.Println("         txt Save data as text")
	fmt.Println("         json Save data as JSON")
	fmt.Println("         Default: txt")
	fmt.Println("      Example:")
	fmt.Println("         cli -SerializeMode=txt")
	fmt.Println("         cli -SerializeMode=json")
	fmt.Println()

	fmt.Println("Command details:")
	fmt.Println()

	fmt.Println("  compress")
	fmt.Println("    Compress an image with a quality value from 1 to 100.")
	fmt.Println("    Example: cli compress")
	fmt.Println()

	fmt.Println("  crop")
	fmt.Println("    Crop an image using width, height, X and Y coordinates.")
	fmt.Println("    Example: cli crop")
	fmt.Println()

	fmt.Println("  filter")
	fmt.Println("    Apply one of the following filters:")
	fmt.Println("      grayscale")
	fmt.Println("      sepia")
	fmt.Println("      invert")
	fmt.Println("    Example: cli filter")
	fmt.Println()

	fmt.Println("  flip")
	fmt.Println("    Flip an image horizontally or vertically.")
	fmt.Println("    Modes:")
	fmt.Println("      horizontal")
	fmt.Println("      vertical")
	fmt.Println("    Example: cli flip")
	fmt.Println()

	fmt.Println("  format")
	fmt.Println("    Convert an image to another format.")
	fmt.Println("    Supported formats:")
	fmt.Println("      jpg")
	fmt.Println("      png")
	fmt.Println("    Example: cli format")
	fmt.Println()

	fmt.Println("  resize")
	fmt.Println("    Resize an image using a custom width and height.")
	fmt.Println("    Example: cli resize")
	fmt.Println()

	fmt.Println("  rotate")
	fmt.Println("    Rotate an image by a specified angle.")
	fmt.Println("    Example: cli rotate")
	fmt.Println()

	fmt.Println("  watermark")
	fmt.Println("    Add an image watermark to the current image.")
	fmt.Println("    Watermark options:")
	fmt.Println("      Opacity: 0-100")
	fmt.Println("      Mode: single / tile")
	fmt.Println("    Example: cli watermark")
	fmt.Println()

	fmt.Println("  path")
	fmt.Println("    Show the current image path and set a new image path.")
	fmt.Println("    Example: cli path")
	fmt.Println()

	fmt.Println("  list")
	fmt.Println("    Show the images saved for the currently logged-in user.")
	fmt.Println("    Example: cli list")
	fmt.Println()

	fmt.Println("  sign")
	fmt.Println("    Sign in or register a user with username and password.")
	fmt.Println("    Example: cli sign")
	fmt.Println()

	fmt.Println("  exit")
	fmt.Println("    Exit the CLI application.")
	fmt.Println("    Example: cli exit")
	fmt.Println()

	fmt.Println("Examples:")
	fmt.Println()
	fmt.Println("  cli -SerializeMode=txt")
	fmt.Println("  cli compress")
	fmt.Println("  cli crop")
	fmt.Println("  cli filter")
	fmt.Println("  cli flip")
	fmt.Println("  cli format")
	fmt.Println("  cli resize")
	fmt.Println("  cli rotate")
	fmt.Println("  cli watermark")
	fmt.Println("  cli path")
	fmt.Println("  cli list")
	fmt.Println("  cli sign")
	fmt.Println("  cli exit")
	fmt.Println()

	fmt.Println("Notes:")
	fmt.Println("  - Login is required for image-processing commands.")
	fmt.Println("  - Set the image path before using image processing commands.")
	fmt.Println("  - Use 'list' to display images saved for the current user.")
	fmt.Println("  - SerializeMode controls the data save/serialization mode.")
	fmt.Println("  - The default SerializeMode is 'txt'.")
	fmt.Println("  - Follow the prompts shown by each command.")
	fmt.Println("  - Use 'cli help' to display this information again.")
	fmt.Println("  - Use 'exit' to close the application.")
	fmt.Println()
}

func FinalizeCommand() {
	DB.SaveImageDB(Authentication.Username)

}

// RunCommand executes a CLI command and handles its associated user interaction.
//
// Commands that modify or process images require an authenticated user.
// When authentication is required and no user is currently logged in,
// RunCommand prompts the user to log in before continuing.
//
// Supported commands include help, path, sign, compress, crop, filter, flip,
// format, resize, rotate, watermark, and exit.
func RunCommand(command string) {
	if command != "sign" && command != "exit" && Authentication == nil {
		fmt.Println("!!! please Login !!!")
		login()

		return
	}

	switch command {
	case "help":
		printHelp()
	case "path":
		fmt.Printf("Current image path %s\n", ImagePath)
		imagePath()
		fmt.Printf("New Image Path %s\n", ImagePath)
	case "sign":
		fmt.Println("Username: ")
		name := Scanner("")

		fmt.Println("Password: ")
		password := Scanner("")

		result := Login.Sign(name, password)

		if result == "" {
			fmt.Println("Sign up failed.")
			return
		}

		fmt.Println("!!! User signed successfully !!!")

	case "compress":
		fmt.Println("Quality (1-100): ")
		quality := Scanner("")
		fmt.Println("compress")
		result := Build.Compress(ImagePath, quality)

		if result == "" {
			fmt.Println("Compression failed.")

			return
		}

		FinalizeCommand()

		fmt.Println("!!! Compression completed successfully !!!")
		fmt.Println("Output:", result)
	case "crop":
		fmt.Println("Crop width: ")
		width := Scanner("")

		fmt.Println("Crop height: ")
		height := Scanner("")

		fmt.Println("Crop X: ")
		x := Scanner("")

		fmt.Println("Crop Y: ")
		y := Scanner("")

		result := Build.Crop(width, height, x, y, ImagePath)

		if result == "" {
			fmt.Println("Crop failed.")

			return
		}

		FinalizeCommand()

		fmt.Println("!!! Crop completed successfully !!!")
		fmt.Println("Result:", result)

	case "filter":
		fmt.Println("Filter (grayscale/sepia/invert): ")
		filter := Scanner("")

		result := Build.ApplyFilter(ImagePath, filter)

		if result == "" {
			fmt.Println("Filter failed.")

			return
		}

		FinalizeCommand()

		fmt.Println("!!! Filter applied successfully !!!")
		fmt.Println("Output:", result)

	case "flip":
		fmt.Println("Mode (horizontal/vertical): ")
		mode := Scanner("")

		result := Build.Flip(ImagePath, mode)

		if result == "" {
			fmt.Println("Flip failed.")

			return
		}

		FinalizeCommand()

		fmt.Println("!!! Flip completed successfully !!!")
		fmt.Println("Output:", result)

	case "format":
		fmt.Println("Output format (jpg/png): ")
		format := Scanner("")

		result := Build.ChangeFormat(ImagePath, format)

		if result == "" {
			fmt.Println("Format conversion failed.")

			return
		}

		FinalizeCommand()

		fmt.Println("!!! Format conversion completed successfully !!!")
		fmt.Println("Output:", result)

	case "resize":
		fmt.Println("Width: ")
		width := Scanner("")

		fmt.Println("Height: ")
		height := Scanner("")

		result := Build.Resize(ImagePath, width, height)

		if result == "" {
			fmt.Println("Resize failed.")

			return
		}

		FinalizeCommand()

		fmt.Println("!!! Resize completed successfully !!!")
		fmt.Println("Output:", result)

	case "rotate":
		fmt.Println("Rotation angle: ")
		rotate := Scanner("")

		result := Build.Rotate(ImagePath, rotate)

		if result == "" {
			fmt.Println("Rotate failed.")

			return
		}

		FinalizeCommand()

		fmt.Println("!!! Rotate completed successfully !!!")
		fmt.Println("Output:", result)

	case "watermark":
		fmt.Println("Watermark image path: ")
		watermarkPath := Scanner("")

		fmt.Println("Opacity (0-100): ")
		opacity := Scanner("")

		fmt.Println("Watermark mode (single/tile): ")
		mode := Scanner("")

		result := Build.Watermark(
			ImagePath,
			watermarkPath,
			opacity,
			mode,
		)

		if result == "" {
			fmt.Println("Watermark failed.")

			return
		}

		FinalizeCommand()

		fmt.Println("!!! Watermark applied successfully !!!")
		fmt.Println("Output:", result)

	case "list":
		DB.LoadImageDB()

		if len(DB.ImageDB) == 0 {
			fmt.Println("No images found.")
			return
		}

		fmt.Println("\nImages:")
		fmt.Println("-------------------------------")

		for _, f := range DB.ImageDB {
			if f.User == Authentication.Username {
				fmt.Printf("Created by %v, File path %v\n", f.User, f.Path)
			}
		}
		fmt.Println("")

	case "exit":
		os.Exit(0)

	}
}
