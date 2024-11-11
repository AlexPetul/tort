import {List, Avatar, notification} from "antd"

import { AddToFavourites, LoadCatalog } from "../wailsjs/go/main/App"
import { useEffect, useState } from "react"

export const Catalog = () => {
    const [data, setData] = useState([])
    const [api, contextHolder] = notification.useNotification();

    useEffect(() => {
        LoadCatalog().then((result) => {
            setData(result)
        });
    }, [])

    const formatStars = (number) => {
        if (number >= 1000) {
            let result = parseFloat(number / 1000).toFixed(1)
            if (result.endsWith("0")) {
                result = result.split(".")[0]
            }
            return result + "k"
        }
        return number
    }

    const addToFavourites = (e) => {
        const id = e.target.getAttribute("data-id")
        const name = e.target.getAttribute("data-name")
        AddToFavourites(parseInt(id)).then(result => {
            api.open({
                message: `Project "${name}" was added to your favourites`,
                duration: 40,
                style: {backgroundColor: '#28a745'}
            })
            const newData = data.map((item) => {
                if (item.id === parseInt(id)) {
                    return { ...item, is_favourite: true };
                }
                return item;
            });
            setData(newData);
    
        });
    }

    return (
        <>
            {contextHolder}
            <List
                itemLayout="horizontal"
                dataSource={data}
                renderItem={(item, index) => (
                    <List.Item>
                        <List.Item.Meta
                            className="catalog-item"
                            avatar={<Avatar src={item.avatar_url} />}
                            title={
                                <>
                                    <a href={item.git_url}>{item.name}</a>
                                    <div className="float-end pe-auto">
                                        {item.is_favourite ?
                                            <i className="fas fa-check text-success"></i>
                                        :
                                            <i
                                                className="fas fa-plus text-primary"
                                                data-id={item.id}
                                                data-name={item.name}
                                                onClick={addToFavourites}
                                            ></i>
                                        }
                                    </div>
                                </>
                            }
                            description={
                                <>
                                    <span><i className="fas fa-star text-warning"></i> {formatStars(item.stars)}</span>
                                    <span> {item.owner}</span>
                                </>
                            }
                        />
                    </List.Item>
                )}
            />
        </>
    )
}