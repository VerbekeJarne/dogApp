import React, { useState, useEffect } from "react";

function App() {
    const [dogs, setDogs] = useState([]);

    useEffect(() => {
        fetch("http://localhost:8080/api/dogs")
            .then((response) => response.json())
            .then((data) => setDogs(data))
            .catch(error => console.log(error));
    }, []);
    return (
        <div>
            <h1>Dogs in the API:</h1>
            <ul>
                {dogs.map(dog => (
                    <li key={dog.name}>
                        <div>Name: {dog.name}</div>
                        <div>Breed: {dog.breed}</div>
                        <div>Colour: {dog.colour}</div>
                    </li>
                ))}
            </ul>
        </div>
    );
}

export default App;