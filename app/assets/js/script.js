window.addEventListener("DOMContentLoaded", () => {
    fetchHorses();
  });
  
  function fetchHorses() {
    fetch("/api/horses")
      .then((response) => response.json())
      .then((horses) => {
        displayHorses(horses);
      });
  }
  
  function displayHorses(horses) {
    const horseContainer = document.getElementById("horse-container");
    horseContainer.innerHTML = "";
  
    horses.forEach((horse) => {
      const horseCard = document.createElement("div");
      horseCard.className = "horse-card";
      horseCard.innerHTML = `
        <h2 class="horse-name">${horse.Name}</h2>
        <p class="horse-odds">Odds: ${horse.Odds}</p>
        < class="horse-delete" onclick="deleteHorse('${horse.ID}')">Delete</button>
      `;
      horseContainer.appendChild(horseCard);
    });
  }
  
  function deleteHorse(id) {
    fetch(`/api/horses/${id}`, {
      method: "DELETE",
    }).then(() => {
      fetchHorses();
    });
  }