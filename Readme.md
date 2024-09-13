# Music Brainz

## Project Overview

Music Brainz is a web application that fetches and manipulates data from a provided API to create a user-friendly website. It displays detailed information about bands and artists, showcasing various data visualizations and implementing client-server interactions to trigger actions.

## Features

- Detailed artist and band information display
- Concert location mapping
- Upcoming and past concert date listings
- Interactive data visualizations
- Client-server interactions for dynamic content updates

## API Structure

The API consists of four main parts:

1. **Artists**: Includes details about the bands or artists, such as:
   - Name(s)
   - Image
   - Year they began activity
   - First album release date
   - Band members

2. **Locations**: Contains the last and/or upcoming concert locations.

3. **Dates**: Contains the last and/or upcoming concert dates.

4. **Relations**: Links the data between artists, dates, and locations.

## Technologies Used

- **Backend**: Go
- **Frontend**: HTML, CSS, JavaScript
- **Data Format**: JSON (for API data manipulation and storage)

## Prerequisites

- Go 1.x or higher
- Web browser (Chrome, Firefox, Safari, etc.)

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/music-brainz.git
   ```
2. Navigate to the project directory:
   ```
   cd music-brainz
   ```
3. Install dependencies:
   ```
   go mod tidy
   ```

## Usage

1. Start the server:
   ```
   go run main.go
   ```
2. Open a web browser and navigate to `http://localhost:8080` (or whatever port you've configured)

## Project Structure

```
music-brainz/
├── main.go
├── api/
│   ├── client.go
├── config/
│   ├── config.go  
├── funcs/
│   ├── func.sgo
├── handlers/
│   ├── aboutUs.go
│   ├── artist.go
│   ├── location.go
│   └── date.go
├── static/
│   ├── css/
│   │   └── styles.css
│   └── templates/
│       ├── about.html
│       ├── index.html
│       └── artist.html
└──README.md
```

## API Endpoints

- `/api/artists`: Get all artists
- `/api/artists/{id}`: Get a specific artist
- `/api/locations`: Get all concert locations
- `/api/dates`: Get all concert dates

## Contributing

Contributions are welcome! Here's how you can contribute:

1. Fork the repository
2. Create a new branch (`git checkout -b feature/AmazingFeature`)
3. Make your changes
4. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
5. Push to the branch (`git push origin feature/AmazingFeature`)
6. Open a Pull Request

Please ensure your code adheres to the existing style and includes appropriate tests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Authors

- **Mehdi** - Project creator and developer

## Acknowledgments

- Thanks to the Music Brainz API for providing the data
- Inspiration from similar music information platforms

## Contact

Mehdi - ig := [@mehdim_dev]
Discord-profile := [https://discord.com/users/720655054834630676]

Project Link: [https://github.com/mehdi/music-brainz](https://github.com/mehdi/music-brainz)
