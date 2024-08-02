import React, {useEffect, useState} from 'react';
import './App.css';
import Login from "./pages/Login";
import Nav from "./components/Nav";
import {BrowserRouter, Route} from "react-router-dom";
import Home from "./pages/Home";
import Register from "./pages/Register";
import WelcomePage from './pages/WelcomePage';


function App() {

    const [ID, setID] = useState('');

    useEffect(() => {
        (
            async () => {
                const response = await fetch('http://localhost:8000/api/user', {
                    headers: {'Content-Type': 'application/json','Access-Control-Allow-Origin': 'http://localhost:3000/','Access-Control-Allow-Credentials':'true', 'Access-Control-Allow-Headers': 'Content-Type, Authorization'},
                    credentials: 'same-origin',
                    mode: 'cors'
                });

                const content = await response.json();
                if(content.message==="unauthenticated"){
                    setID('');
                }else{
                    setID(content.info.id);
                }
                

            }
        )();
    });


    return (
        <div className="App">
            <BrowserRouter>
                <Nav ID={ID} setID={setID}/>

                <main className="form-signin">
                    <Route path="/" exact component={Home}/>
                    <Route path="/login" component={Login}/>
                    <Route path="/register" component={Register}/>
                    <Route path="/WelcomePage" component={WelcomePage}/>
                </main>
            </BrowserRouter>
        </div>
    );
}

export default App;
