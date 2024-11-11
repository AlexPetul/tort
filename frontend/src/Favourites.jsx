import {List, Avatar} from "antd"

import { AlignTags, DeleteFavouriteItem, LoadFavourites } from "../wailsjs/go/main/App"
import { useEffect, useState } from "react"

export const Favourites = () => {
    const [data, setData] = useState([])

    useEffect(() => {
        LoadFavourites().then((result) => {
            setData(result)
        });
    }, [])

    const deleteItem = (e) => {
        const id = e.target.getAttribute("data-id")
        DeleteFavouriteItem(parseInt(id)).then(response => {
            const newData = data.filter((item) => {
                if (item.id !== parseInt(id)) {
                    return item
                }
            });
            setData(newData);
        })
    }

    const upgradeTag = (e) => {
        const id = e.target.getAttribute("data-id")
        AlignTags(parseInt(id)).then(response => {
            const newData = data.map((item) => {
                if (item.id === parseInt(id)) {
                    return { ...item, current_release: item.latest_release };
                }
                return item;
            });
            setData(newData);
        })
    }

    return (
        <List
            itemLayout="horizontal"
            dataSource={data}
            renderItem={(item, index) => (
                <List.Item>
                    <List.Item.Meta
                        className="catalog-item"
                        avatar={<Avatar src={item.catalog_item.avatar_url} />}
                        title={
                            <>
                                <div>
                                    {item.catalog_item.name}&nbsp;
                                    {item.current_release !== item.latest_release 
                                    ?
                                     <i 
                                        className="fas fa-fire-flame-curved text-warning pe-auto"
                                        data-id={item.id}
                                        onClick={upgradeTag}
                                     />
                                    : null
                                    }
                                </div>
                                <div className="float-end pe-auto">
                                    <i 
                                        className="fas fa-trash-can text-danger"
                                        data-id={item.id}
                                        onClick={deleteItem}
                                    ></i>
                                </div>
                            </>
                        }
                        description={
                            <>
                                <span>Current: {item.current_release}</span><br/>
                                <span>Latest: {item.latest_release}</span>
                            </>
                        }
                    />
                </List.Item>
            )}
        />
    )
}