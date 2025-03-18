/**
 * Main JavaScript file for {{.ProjectName}}
 * Author: {{.Author}}
 * Year: {{.Year}}
 */

document.addEventListener('DOMContentLoaded', function() {
    console.log('{{.ProjectName}} application initialized');
    
    // Add event listeners and other initialization code here
    document.querySelectorAll('a').forEach(link => {
        link.addEventListener('click', function(e) {
            console.log('Link clicked:', this.href);
        });
    });
    
    // Example: Display a welcome message after a short delay
    setTimeout(function() {
        const mainElement = document.querySelector('main');
        if (mainElement) {
            const welcomeElement = document.createElement('div');
            welcomeElement.className = 'welcome-message';
            welcomeElement.textContent = 'Welcome to {{.ProjectName}}!';
            welcomeElement.style.padding = '1rem';
            welcomeElement.style.marginTop = '1rem';
            welcomeElement.style.backgroundColor = '#e3f2fd';
            welcomeElement.style.borderRadius = '5px';
            welcomeElement.style.textAlign = 'center';
            mainElement.appendChild(welcomeElement);
        }
    }, 1000);
}); 