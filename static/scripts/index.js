let isScrolling;

window.addEventListener("scroll", function () {
  // Clear the timeout if it exists
  clearTimeout(isScrolling);

  // Add a class to indicate that the user is scrolling
  document.body.classList.add("scrolling");

  // Set a timeout to run after scrolling ends
  isScrolling = setTimeout(function () {
    // Remove the class after scrolling ends
    document.body.classList.remove("scrolling");
  }, 100); // Adjust the timeout duration as needed
});
