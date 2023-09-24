import React from "react"
import '../styles/sidebar.css'
import CreateChatIcon from '../images/icons8-new-message-90.png'
import FriendsIcon from '../images/icons8-users-90.png'

const Sidebar = () => {
    return (
        <div className="sidebar">
            <div className="sidebar-header">
                <div className="profile-section">
                    <div className="profile-picture"></div>
                </div>
                <div className="icon-section">
                    <div className="icon-container">
                        <img role="button" src={FriendsIcon} alt='Create-Chat-Icon' className="header-icon"/>
                    </div>
                    <div className="icon-container">
                        <img role='button' src={CreateChatIcon} alt='Create-Chat-Icon' className="header-icon"/>
                    </div>
                </div>
            </div>
            <div className="sidebar-context">
                <div className="chat-bar">
                    <div className="chat-bar-container">
                        <div className="friend-picture-container">
                            
                        </div>
                        <div className="username-container">
                            <p>enes furkan olcay</p>
                        </div>
                    </div>
                </div>
                <div className="chat-bar">
                    <div className="chat-bar-container">
                        <div className="friend-picture-container">
                            
                        </div>
                    </div>
                </div>
                <div className="chat-bar">
                    <div className="chat-bar-container">
                        <div className="friend-picture-container">
                            
                        </div>
                    </div>
                </div>
                <div className="chat-bar">
                    <div className="chat-bar-container">
                        <div className="friend-picture-container">
                            
                        </div>
                    </div>
                </div>
                <div className="chat-bar">
                    <div className="chat-bar-container">
                        <div className="friend-picture-container">
                            
                        </div>
                    </div>
                </div>
                <div className="chat-bar">
                    <div className="chat-bar-container">
                        <div className="friend-picture-container">
                            
                        </div>
                    </div>
                </div>
                <div className="chat-bar">
                    <div className="chat-bar-container">
                        <div className="friend-picture-container">
                            
                        </div>
                    </div>
                </div>
                <div className="chat-bar">
                    <div className="chat-bar-container">
                        <div className="friend-picture-container">
                            
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default Sidebar
