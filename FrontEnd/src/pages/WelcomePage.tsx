
const WelcomePage = (props:{data?: Object}) => {

    const values = JSON.parse(JSON.stringify(props.data))
    return (    
        <div> 
            {'Welcome to the portal, '+values.name+ ' You have '+values.role+' role and have '+values.Permissons+' permisssion(s) for '+ values.Departments+' department(s)'}
        </div>
        
    );
};

export default WelcomePage