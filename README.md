# Image Processing System (IPS)

IPS is an image processing system that allows users to manipulate and transform images through an interactive CLI.

The project provides a variety of image processing features, including compression, cropping, filtering, flipping, format conversion, resizing, rotation, and watermarking.

## Current Status

The project is currently being developed in CLI mode.

Server/API support was previously part of the project and will be reintroduced and further developed in a future version.

## License & Third-Party Dependencies

This project uses several open-source packages. The packages and their respective licenses are listed below.

### 1. MySQL Driver

**Package:** `go-sql-driver/mysql`
**License:** Mozilla Public License 2.0 (MPL-2.0)

[License](https://github.com/go-sql-driver/mysql/blob/master/LICENSE)

### 2. Gin

**Package:** `gin-gonic/gin`
**License:** MIT License

[License](https://github.com/gin-gonic/gin?tab=MIT-1-ov-file)

### 3. UUID

**Package:** `google/uuid`
**License:** BSD 3-Clause License

[License](https://github.com/google/uuid?tab=License-1-ov-file)

### 4. Imaging

**Package:** `disintegration/imaging`
**License:** MIT License

[License](https://github.com/disintegration/imaging?tab=MIT-1-ov-file)

### 5. JWT

**Package:** `golang-jwt/jwt`
**License:** MIT License

[License](https://github.com/golang-jwt/jwt?tab=MIT-1-ov-file)

---

Now that you have an overview of the project and its dependencies, let's take a look at how to use IPS.


# Image Processing System (IPS)

IPS is an image processing system that provides image manipulation features through an interactive CLI.

## Current Status

The project is currently being developed in CLI mode.

Server/API support was previously part of the project and will be reintroduced and further developed in a future version.

## Features

### Image Processing

* **Compress** — Compress an image and reduce its file size.
* **Crop** — Crop an image using custom dimensions and coordinates.
* **Filter** — Apply grayscale, sepia, or invert filters.
* **Flip** — Flip an image horizontally or vertically.
* **Change Format** — Convert an image between supported formats.
* **Resize** — Resize an image using a custom width and height.
* **Rotate** — Rotate an image by a specified angle.
* **Watermark** — Add an image watermark with configurable opacity and display mode.

### CLI Features

* **Path** — View and change the current image path.
* **Sign** — Sign in or register a user.
* **Help** — Display available commands and usage information.
* **Exit** — Close the CLI application.

## Available Commands

| Command     | Description                                               |
|-------------|-----------------------------------------------------------|
| `compress`  | Compress an image and reduce its file size                |
| `crop`      | Crop an image to the specified dimensions and position    |
| `filter`    | Apply a filter to an image                                |
| `flip`      | Flip an image horizontally or vertically                  |
| `format`    | Convert an image to another format                        |
| `resize`    | Resize an image to the specified dimensions               |
| `rotate`    | Rotate an image by a specified angle                      |
| `watermark` | Add an image watermark with configurable opacity and mode |
| `path`      | Show or change the current image path                     |
| `sign`      | Sign in or register a user                                |
| `help`      | Show help information                                     |
| `exit`      | Exit the application                                      |

## Image Processing Details

### Compress

Compress an image using a quality value from `1` to `100`.

```text
cli compress
```

The CLI will prompt for the required information.

### Crop

Crop an image using:

* Width
* Height
* X coordinate
* Y coordinate

```text
cli crop
```

### Filter

Apply one of the supported filters:

* `grayscale`
* `sepia`
* `invert`

```text
cli filter
```

### Flip

Flip the image using one of the following modes:

* `horizontal`
* `vertical`

```text
cli flip
```

### Change Format

Convert an image to another supported format.

Currently supported formats:

* `jpg`
* `png`

```text
cli format
```

### Resize

Resize an image using a custom width and height.

```text
cli resize
```

### Rotate

Rotate an image by a specified angle.

```text
cli rotate
```

### Watermark

Add an image watermark to the current image.

Watermark options include:

* **Opacity:** `0–100`
* **Mode:** `single` or `tile`

```text
cli watermark
```

## Path Management

The `path` command is used to view or change the current image path.

```text
cli path
```

Set the image path before using image-processing commands.

## Authentication

The CLI includes user authentication through the `sign` command.

Users can:

* Sign in with an existing account.
* Register a new account.

```text
cli sign
```

Authentication is required before using image-processing commands.

## Help

Display the available commands and their descriptions:

```text
cli help
```

## Exit

Close the CLI application:

```text
cli exit
```

You can also use:

```text
exit
```

## Usage Notes

* Sign in before using image-processing commands.
* Set the image path before processing an image.
* Follow the prompts displayed by each command.
* Use `cli help` to view the available commands.
* Use `exit` to close the application.

## Example Workflow

A typical workflow looks like this:

```text
cli sign
cli path
cli resize
cli filter
cli compress
cli watermark
```

Each command guides the user through the required options.

## Project Roadmap

The current focus of the project is the CLI implementation.

Future development will include bringing back the server/API layer and expanding the system with additional features.
