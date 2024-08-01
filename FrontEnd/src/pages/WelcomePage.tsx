// interface LoginResponse {
//     id: number;
//     name: string;
//     email: string;
//     role: string;
//     Permissions: string[];
//     Departments: string[];
// }
const WelcomePage = (props:{data?: Object}) => {

    // content :JSON= props.data?props.data.info: 

//     content: Object = props.data;
//    content = JSON.parse(content||"")
    
    const values = JSON.parse(JSON.stringify(props.data))

    console.log("entered",values.id);
    return (    

        <div>
            
            {'Welcome to the portal, '+values.name+ ' You have '+values.role+' role and have '+values.Permissons+' permisssion(s) for '+ values.Departments+' department(s)'}
        </div>
    );
};

export default WelcomePage