

document.addEventListener('DOMContentLoaded', function () 
    {
        fetch('/api/recipes')
            .then(response => response.json())
            .then(recipes => {
                const selectElement = document.getElementById('recipe-select');
                recipes.forEach(recipe => {
                    const option = new Option(recipe.title, recipe.ID);
                    selectElement.add(option);
                });
            })
            .catch(error => console.error('Error:', error));
    fetchRecipes();
    attachEventListeners();
    populateRecipeDropdown();
    deleteRecipe();
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
    
    
// Get the editForm element and the current recipe ID to edit
const editForm = document.querySelector('#form-update-recipe');
const recipeToEdit = editForm && editForm.dataset.recipeid;

// Add an event listener to listen for the form submit
editForm && editForm.addEventListener('submit', (event) => {
    // Prevent the default behaviour of the form element
    event.preventDefault();

    // Convert form data into a JavaScript object
    const formData = Object.fromEntries(new FormData(editForm));

    return fetch(`/recipe/${recipeToEdit}`, {
        // Use the PATCH or PUT method, depending on your API
        method: 'PATCH',
        headers: {
            'Content-Type': 'application/json'
        },
        // Convert the form's Object data into JSON
        body: JSON.stringify(formData),
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        // Redirect to the updated recipe's detail view, or handle according to your app's logic
        document.location.href = `/recipe/${recipeToEdit}`;
    })
    .catch((error) => {
        console.error('There has been a problem with your fetch operation:', error);
        // Handle errors here, such as displaying a message to the user
    });
});



    
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

    function handleFormSubmitWithFetch(url, method) {
        return function(event) {
            event.preventDefault();
            const formData = new FormData(event.target);
    
            fetch(url, { method: method, body: formData })
                .then(response => {
                    if (!response.ok) {
                        throw new Error(`HTTP error! status: ${response.status}`);
                    }
                    return response.text();
                })
                .then(data => {
                    console.log('Operation successful:', data);
                    window.location.href = '/recipes'; // Redirect or update the UI as needed
                })
                .catch(error => console.error('Error:', error));
        };
        
    }