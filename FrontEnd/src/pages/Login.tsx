import React, {SyntheticEvent, useState} from 'react';
import {Redirect} from "react-router-dom";
import WelcomePage from './WelcomePage';


const Login =()=> {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [redirect, setRedirect] = useState(false);
    const [loginFailed, setLoginFailed] = useState(false);

  

    const [data, setData] = useState({})

    const submit = async (e: SyntheticEvent) => {
        e.preventDefault();

        const response = await fetch('http://localhost:8000/api/login', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            credentials: 'include',
            body: JSON.stringify({
                email,
                password
            })
        });

      

        const content = await response.json();


        setData(content.info)

        console.log(data)

        console.log(typeof content)

        if(content.message==="success"){
            console.log('redirected to welcome page');
            setRedirect(true);
        }else if(content.message.id === 0){
            console.log("user not found");
            setLoginFailed(true)
        }
        
       
        // props.setRole(content.role);
        // props.setPermissions(content.permissions);
        // props.setDepartment(content.department)

    }

    if (redirect) {
        console.log('going to welcomePage', data);
       return <WelcomePage data={data} /> ; // changes made on 31st of july
    }
    if(loginFailed){
        return <Redirect to="/Register"/>
    }
    

    return (
        <form onSubmit={submit}>
            <h1 className="h3 mb-3 fw-normal">Login</h1>
            <input type="email" className="form-control" placeholder="Email address" required
                   onChange={e => setEmail(e.target.value)}
            />

            <input type="password" className="form-control" placeholder="Password" required
                   onChange={e => setPassword(e.target.value)}
            />

            <button className="w-100 btn btn-lg btn-primary" type="submit">Sign in</button>
        </form>
    );
};

export default Login;
