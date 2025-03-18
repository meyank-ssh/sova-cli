/* 
   Main stylesheet for {{.ProjectName}} 
   Author: {{.Author}}
   Year: {{.Year}}
*/

:root {
    --primary-color: #3498db;
    --secondary-color: #2ecc71;
    --dark-color: #2c3e50;
    --light-color: #ecf0f1;
    --font-main: 'Arial', sans-serif;
}

* {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
}

body {
    font-family: var(--font-main);
    line-height: 1.6;
    color: var(--dark-color);
    background-color: var(--light-color);
}

header {
    background-color: var(--primary-color);
    color: white;
    text-align: center;
    padding: 2rem 1rem;
}

header h1 {
    margin-bottom: 0.5rem;
}

main {
    max-width: 1200px;
    margin: 0 auto;
    padding: 2rem;
}

section {
    background-color: white;
    border-radius: 5px;
    padding: 2rem;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}

h2 {
    color: var(--primary-color);
    margin-bottom: 1rem;
}

footer {
    text-align: center;
    padding: 1rem;
    background-color: var(--dark-color);
    color: white;
    position: absolute;
    bottom: 0;
    width: 100%;
} 