document.addEventListener('DOMContentLoaded', function () {
    fetchRecipes();

    const ftbr = document.getElementById("fltrtbr");
    ftbr.addEventListener("click", () => filterRecipes('not-tried'));

    const fcook = document.getElementById("fltrcook");
    fcook.addEventListener("click", () => filterRecipes('cooked'));

    const ffav = document.getElementById("fltrfav");
    ffav.addEventListener("click", () => filterRecipes('favorite'));

    const ftried = document.getElementById("fltrtried");
    ftried.addEventListener("click", () => filterRecipes('tried'));

    document.getElementById('recipeList').addEventListener('click', (event) => {
        const clickedRecipeElement = event.target.closest('.recipe-list-container');
        if (!clickedRecipeElement) return;

        // Read recipe ID from the clicked recipe element
        const recipeID = clickedRecipeElement.dataset.recipeId;

        // Handle "Mark As" button clicks
        const markAsButton = event.target.closest('.mark-as-button');
        if (markAsButton) {
            const status = markAsButton.dataset.status;
            updateRecipeStatus(recipeID, status);
        }

        // Handle "Delete" button clicks
        const deleteButton = event.target.closest('.delete-button');
        if (deleteButton) {
            deleteRecipe(recipeID);
        }
    });

    function fetchRecipes() {
        // Fetch the initial HTML content
        fetch('/')
            .then(response => {
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                return response.text(); // Parse the HTML content
            })
            .then(htmlContent => {
                // Create a temporary container to parse the HTML
                const tempContainer = document.createElement('div');
                tempContainer.innerHTML = htmlContent;

                // Extract the recipe list container from the parsed HTML
                const recipeListContainer = tempContainer.querySelector('#recipeList');

                if (recipeListContainer) {
                    // Update the existing recipe list container with the received HTML content
                    const currentRecipeListContainer = document.getElementById('recipeList');
                    currentRecipeListContainer.innerHTML = recipeListContainer.innerHTML;

                    // Add click event listeners to each recipe element
                    currentRecipeListContainer.addEventListener('click', (event) => {
                        const recipeElement = event.target.closest('.recipe-list-container');
                        if (recipeElement) {
                            const recipeID = recipeElement.dataset.recipeId;
                            showRecipeDetails(recipeID);
                        }
                    });
                } else {
                    console.error('Recipe list container not found in the received HTML.');
                }
            })
            .catch(error => console.error('Error:', error));
    }

    function filterRecipes(status) {
        var recipeListContainer = document.getElementById("recipeList");
        // Make a GET request to the server to fetch the list of recipes with the given status
        fetch(`/api/recipes/${status}`)
            .then(response => {
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                return response.json();
            })
            .then(recipes => {

                // Clear the existing recipe list
                recipeListContainer.innerHTML = '';

                // Add each recipe to the recipeListContainer div
                recipes.forEach(recipe => {
                    var recipeElement = document.createElement('div');
                    recipeElement.innerHTML = `
                        <p>${recipe.title}: ${recipe.status}</p>
                    `;

                    recipeListContainer.appendChild(recipeElement);
                });
            })
            .catch(error => console.error('Error:', error));
    }

    // Handle click events to display recipe details
    function showRecipeDetails(recipeID) {
        fetch(`/api/recipes/${recipeID}`)
            .then(response => response.json())
            .then(recipe => {
                const recipeDetailsContainer = document.getElementById('recipeDetailsContainer');
                const titleElement = document.getElementById('recipeTitle');
                const ingredientsElement = document.getElementById('recipeIngredients');
                const instructionsElement = document.getElementById('recipeInstructions');
                const statusElement = document.getElementById('recipeStatus');
                const idElement = document.getElementById('recipeID');

                titleElement.textContent = recipe.title;
                ingredientsElement.textContent = `Ingredients: ${recipe.ingredients}`;
                instructionsElement.textContent = `Instructions: ${recipe.instructions}`;
                statusElement.textContent = `Status: ${recipe.status}`;
                idElement.textContent = recipe.id;

                // Show the details container
                recipeDetailsContainer.style.display = 'block';
            })
            .catch(error => console.error('Error fetching recipe details:', error));
    }

    // Function to update recipe status
    async function updateRecipeStatus(recipeID, status) {

        const response = await fetch(`/api/recipes/${recipeID}/${status}`, { method: "POST" });

        if (response.ok) {
            // Update the UI or perform additional actions if needed
            console.log(`Recipe status updated to ${status}`);
        } else {
            console.error("Failed to update recipe status");
        }
    }

    // Function to delete a recipe
    function deleteRecipe(recipeID) {
        // Make a DELETE request to the server
        fetch(`/api/recipes/${recipeID}`, {
            method: 'DELETE'
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                // Don't need to parse response as JSON in this case
                return response.text();
            })
            .then(() => {
                // Refresh the recipe list
                fetchRecipes();
            })
            .catch(error => console.error('Error:', error));
    }

    const newForm = document.getElementById('new-form');

    newForm.addEventListener('submit', function (event) {
        event.preventDefault();

        // Create a FormData object from the form
        const formData = new FormData(newForm);

        // Make a POST request to the server
        fetch('/recipe', {
            method: 'POST',
            body: formData,
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                return response.json();
            })
            .then(data => {
                // Handle the response data, if needed
                console.log('Recipe added successfully:', data);

                window.location.href = '/recipes'; // Redirect to the recipe list page
            })
            .catch(error => console.error('Error:', error));
    });
});
