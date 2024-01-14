// ui/static/js/navigation.js
document.addEventListener("DOMContentLoaded", function () {
    // Get the current page URL
    const currentPath = window.location.pathname;

    // Get the navigation links
    const homeLink = document.getElementById("homeLink");
    const aboutLink = document.getElementById("aboutLink");
    const contactsLink = document.getElementById("contactsLink");

    // Show or hide navigation links based on the current page
    if (currentPath === "/") {
        aboutLink.style.display = "inline";
        contactsLink.style.display = "inline";
        homeLink.style.display = "none";
    } else if (currentPath === "/about") {
        aboutLink.style.display = "none";
        contactsLink.style.display = "inline";
        homeLink.style.display = "inline";
    } else if (currentPath === "/contacts") {
        aboutLink.style.display = "inline";
        contactsLink.style.display = "none";
        homeLink.style.display = "inline";
    } else {
        // If none of the above conditions match, show all links
        aboutLink.style.display = "inline";
        contactsLink.style.display = "inline";
        homeLink.style.display = "inline";
    }
});
