# launch-scheduler

Welcome to SpaceTrouble's Launch Scheduler Project!

It's 2049, and SpaceTrouble is revolutionizing space travel by offering regular flights to destinations across our solar system. As a key member of our engineering team, you'll be developing a critical booking system that manages our launch operations alongside our competitor SpaceX.

## Project Overview

We operate multiple launchpads shared with SpaceX, requiring careful scheduling coordination. Our spaceships service an impressive range of destinations including:

- Mars
- Moon
- Pluto
- Asteroid Belt
- Europa
- Titan
- Ganymede

To maximize variety, each launchpad rotates through different destinations daily, ensuring no destination repeats within the same week from a single launchpad.

## Technical Integration

The system will integrate with the SpaceX API (https://github.com/r-spacex/SpaceX-API) to access launchpad availability and SpaceX's launch schedule.

## Required API Endpoints

1. Booking Creation Endpoint
   Create a new booking by accepting passenger details:

   - First Name
   - Last Name
   - Gender
   - Birthday
   - Launchpad ID
   - Destination ID
   - Launch Date

   The endpoint must validate launch availability against SpaceX's schedule to prevent conflicts.

2. Booking Retrieval Endpoint
   Fetch all existing bookings in the system.

## Bonus Features

Enhance your solution with:

- Docker/docker-compose deployment configuration
- Comprehensive unit and functional test coverage
- Booking cancellation endpoint

## Development Requirements

- Backend: Golang
- Database: PostgreSQL
- Version Control: GitHub/Bitbucket with regular, atomic commits

Let's build the future of space travel together! 🚀
