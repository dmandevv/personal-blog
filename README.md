# Personal Blog

A lightweight, self-hosted personal blog application built with Go. It provides a simple way to write, publish, and manage blog articles with a guest-facing interface and a secure admin panel.

This project is an exercise from the roadmap.sh project: [https://roadmap.sh/projects/unit-converter](https://roadmap.sh/projects/personal-blog) 📚

Demo: https://personal-blog-81li.onrender.com

## Features

- **Guest Section**: Browse and read published articles
- **Admin Section**: Create, edit, publish, and delete articles
- **Basic Authentication**: Secure admin access with username and password
- **Article Management**: Full CRUD operations for articles
- **JSON-based Storage**: Articles stored as JSON files for simplicity
- **Environment Configuration**: Easy setup via environment variables and config file

## Project Structure

```
├── main.go               # Application entry point and route setup
├── article.go            # Article data structure
├── auth.go               # Authentication middleware
├── handleHome.go         # Home page handler
├── handleArticle.go      # Article viewing handler
├── handleAdmin.go        # Admin dashboard handler
├── handleNew.go          # New article creation form
├── handlePublish.go      # Article publishing handler
├── handleEdit.go         # Article editing form
├── handleUpdate.go       # Article update handler
├── handleDelete.go       # Article deletion handler
├── json.go               # JSON serialization utilities
├── config.json           # Application configuration
├── static/               # Frontend assets
│   ├── index.html        # Home page
│   ├── article.html      # Article display page
│   ├── admin.html        # Admin dashboard
│   ├── new.html          # Create article form
│   ├── edit.html         # Edit article form
│   └── css/style.css     # Stylesheet
├── articles/             # Stored article JSON files
└── go.mod                # Go module dependencies
```

## Prerequisites

- Go 1.22.2 or later

## Installation

1. **Clone the repository** (or navigate to your project directory):
   ```bash
   git clone https://github.com/dmandevv/personal-blog.git
   cd personal-blog
   ```

2. **Install dependencies**:
   ```bash
   go mod download
   ```

3. **Set up environment variables**:
   
   Create a `.env` file in the project root with the following variables:
   ```
   ARTICLE_DIRECTORY=./articles
   PORT=8080
   ADMIN_USERNAME=your_username
   ADMIN_PASSWORD=your_password
   ADMIN_REALM=YourBlogName
   ```

   Or set them as environment variables:
   ```bash
   export ARTICLE_DIRECTORY=./articles
   export PORT=8080
   export ADMIN_USERNAME=your_username
   export ADMIN_PASSWORD=your_password
   export ADMIN_REALM=YourBlogName
   ```

## Running the Application

```bash
go run .
```

The application will start and listen on the configured port (default: 8080).

Access the blog at: `http://localhost:8080`

## Usage

### Guest Section

- **Home Page** (`/home`): Browse all published articles
- **View Article** (`/article/{id}`): Read a specific article

### Admin Section

Access the admin panel with your configured credentials:

- **Admin Dashboard** (`/admin`): Overview of all articles with actions
- **Create New Article** (`/new`): Write a new article
- **Edit Article** (`/edit/{id}`): Modify an existing article
- **Publish Article** (`/publish`): Publish a draft article
- **Delete Article** (`/delete/{id}`): Remove an article
- **Update Article** (`/update`): Save changes to an article

All admin endpoints are protected by HTTP Basic Authentication.

## Configuration

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `ARTICLE_DIRECTORY` | Path to store article JSON files | `./articles` |
| `PORT` | Port to run the application on | `8080` |
| `ADMIN_USERNAME` | Username for admin authentication | Required |
| `ADMIN_PASSWORD` | Password for admin authentication | Required |
| `ADMIN_REALM` | Realm name for authentication prompt | Optional |

### config.json

The application automatically manages a `config.json` file containing:
- `ArticleDirectory`: Path to articles folder
- `Port`: Server port
- `next_article_id`: Auto-incrementing ID for new articles

This file is generated/updated automatically on startup.

## Article Structure

Articles are stored as JSON files in the articles directory with the following structure:

```json
{
  "id": 1,
  "title": "My First Article",
  "content": "Article content here...",
  "date_published": "2026-01-16T10:30:00Z"
}
```

## Security

- Admin endpoints use HTTP Basic Authentication with SHA-256 hashing
- Credentials are never stored in the application
- Use strong, unique passwords for admin access
- Consider running behind HTTPS in production

## Building for Production

```bash
go build -o personal-blog .
```

Then run the compiled binary:
```bash
./personal-blog
```

## Development

- **Static Files**: Located in the `static/` directory; modify HTML/CSS as needed
- **Handlers**: Each route has its own handler file for easy maintenance
- **Routes**: All routes are defined in `main.go`

## License

MIT License. See the [LICENSE](LICENSE) file for information.
