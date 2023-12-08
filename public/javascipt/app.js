  // public/javascript/app.js
document.addEventListener('DOMContentLoaded', function () {
    fetchRecipes();

    const cookedBtn = document.getElementById("fltrcooked");
    cookedBtn.addEventListener("click", () => filterRecipes('cooked'));

    const cookBtn = document.getElementById("fltrcook");
    cookBtn.addEventListener("click", () => filterRecipes('to-cook'));

    const triedBtn = document.getElementById("fltrtried");
    triedBtn.addEventListener("click", () => filterRecipes('tried'));

    const notTriedBtn = document.getElementById("fltrnottried");
    notTriedBtn.addEventListener("click", () => filterRecipes('not-tried'));

    document.getElementById('recipeList').addEventListener('click', (event) => {
        const clickedRecipeElement = event.target.closest('.recipe-list-container');
        if (!clickedRecipeElement) return;

        const recipeID = clickedRecipeElement.dataset.recipeId;

        const markAsButton = event.target.closest('.mark-as-button');
        if (markAsButton) {
            const status = markAsButton.dataset.status;
            updateRecipeStatus(recipeID, status);
        }

        const deleteButton = event.target.closest('.delete-button');
        if (deleteButton) {
            deleteRecipe(recipeID);
        }
    });

    function fetchRecipes() {
        fetch('/api/recipes')
            .then(response => {
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                return response.json();
            })
            .then(recipes => {
                const recipeListContainer = document.getElementById('recipeList');
                recipeListContainer.innerHTML = '';

                recipes.forEach(recipe => {
                    const recipeElement = document.createElement('div');
                    recipeElement.innerHTML = `
                        <p>${recipe.title} - ${recipe.status}</p>
                    `;

                    recipeListContainer.appendChild(recipeElement);
                });
            })
            .catch(error => console.error('Error:', error));
    }

    function filterRecipes(status) {
        const recipeListContainer = document.getElementById("recipeList");

        fetch(`/api/${status}-recipes`)
            .then(response => {
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                return response.json();
            })
            .then(recipes => {
                recipeListContainer.innerHTML = '';

                recipes.forEach(recipe => {
                    const recipeElement = document.createElement('div');
                    recipeElement.innerHTML = `
                        <p>${recipe.title} - ${recipe.status}</p>
                    `;

                    recipeListContainer.appendChild(recipeElement);
                });
            })
            .catch(error => console.error('Error:', error));
    }

    function updateRecipeStatus(recipeID, status) {
        fetch(`/api/recipes/${recipeID}/${status}`, { method: "POST" })
            .then(response => {
                if (response.ok) {
                    console.log(`Recipe status updated to ${status}`);
                } else {
                    console.error("Failed to update recipe status");
                }
            })
            .catch(error => console.error('Error:', error));
    }

    function deleteRecipe(recipeID) {
        fetch(`/api/recipes/${recipeID}`, {
            method: 'DELETE'
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                return response.text();
            })
            .then(() => {
                fetchRecipes();
            })
            .catch(error => console.error('Error:', error));
    }

    const newForm = document.getElementById('new-form');

    newForm.addEventListener('submit', function (event) {
        event.preventDefault();
    
        const formData = new FormData(newForm);
        const jsonData = {};
    
        formData.forEach((value, key) => {
            jsonData[key] = value;
        });
    
        fetch('/recipe', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(jsonData),
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                return response.json();
            })
            .then(data => {
                console.log('Recipe added successfully:', data);
                window.location.href = '/recipes';
            })
            .catch(error => console.error('Error:', error));
    });
    
});
