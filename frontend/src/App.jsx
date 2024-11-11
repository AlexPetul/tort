import {useState} from 'react';
import { Catalog } from './Catalog';
import './App.css';
import 'bootstrap/dist/css/bootstrap.css';
import { Favourites } from './Favourites';

function App() {
    const [activeTab, setActiveTab] = useState('tab-1')
    
    const onTabClick = (e) => {
        setActiveTab(e.target.getAttribute('data-tab'))
    }

    const renderContent = () => {
        if (activeTab === 'tab-1') {
            return <Favourites/>
        } else if (activeTab === 'tab-2') {
            return <Catalog/>
        }
    }

    return (
        <div id="App">
            <div className="tabs-container">
                <div className="tabs">
                    <div className="tab-links">
                        <button onClick={onTabClick} className="tab-link" data-is-active={activeTab === 'tab-1'} data-tab="tab-1">
                            <i className="fas fa-eye"></i> Favourites
                        </button>
                        <button onClick={onTabClick} className="tab-link" data-is-active={activeTab === 'tab-2'} data-tab="tab-2">
                            <i className="fas fa-list"></i> Catalog
                        </button>
                    </div>
                </div>
            </div>

            {renderContent()}
            
        </div>
    )
}

export default App
