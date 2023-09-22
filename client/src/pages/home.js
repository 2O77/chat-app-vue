import React from "react";
import Main from '../components/main.js'
import Sidebar from "../components/sidebar.js";

const Home = () => {
    return (
        <div className="home">
            <Sidebar/>
            <Main/>
        </div>
    )
}

export default Home;
