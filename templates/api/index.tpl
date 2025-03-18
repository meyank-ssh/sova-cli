<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.ProjectName}}</title>
    <link rel="stylesheet" href="/static/css/style.css">
</head>
<body>
    <header>
        <h1>{{.ProjectName}}</h1>
        <p>{{.ProjectDescription}}</p>
    </header>
    
    <main>
        <section>
            <h2>Welcome to your new web application!</h2>
            <p>This is a starter template for your Go web application.</p>
        </section>
    </main>
    
    <footer>
        <p>&copy; {{.Year}} {{.Author}}. All rights reserved.</p>
    </footer>
    
    <script src="/static/js/app.js"></script>
</body>
</html> 