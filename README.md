# CineLabeler

CineLabeler is a command-line tool designed to automatically rename movie files based on metadata fetched from movie databases. It supports integration with both OMDB and TMDB APIs to retrieve accurate movie information.

## ⚡ Getting Started

### Prerequisites

Before you start using CineLabeler, ensure you have Go installed on your system. Additionally, you will need API keys for OMDB and TMDB if you want to fetch data from these sources.

### Installation

Build the project using the Go command:

```bash
go build -o cineLabeler
```

### Configuration

Set up the required environment variables for OMDB and TMDB API keys, and your preferred language for movie data:

```bash
export OMDB_API_KEY='your_omdb_api_key'
export TMDB_API_KEY='your_tmdb_api_key'
export LANG='fr-FR' # Set to your preferred language, e.g. 'en-US' or 'fr-FR' https://developer.themoviedb.org/docs/languages
```

### Usage

To use CineLabeler, simply run the executable with the path to the movie file you wish to rename:

```bash
./cineLabeler <path_to_original_file.mkv>
```

The tool will automatically fetch the movie's metadata and rename the file accordingly.

## 🍰 Contributing
Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

## ❤️ Support
A simple star to this project repo is enough to keep me motivated on this project for days. If you find your self very much excited with this project let me know with a tweet.

If you have any questions, feel free to reach out to me on [Twitter](https://twitter.com/xxPHDxx).