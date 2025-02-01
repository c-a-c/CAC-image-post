/**
 * Homeに表示する画像一覧を管理する
 * Firebaseに登録された、最新の履歴の画像を上部に表示するようにしたい
 */

import React, {useState, useEffect} from "react";
import Image from "../../types/image";
import {formatTimestamp} from "../../utils/formatTimestamp";
import {ref, getDownloadURL, StorageReference} from "firebase/storage";
import {storage} from "../../lib/firebase";
import "./PreviewImage.css";

const PreviewImage: React.FC<{ image: Image }> = ({image}) => {
    const [prevUrl, setPrevUrl] = useState<string>("");
    const [error, setError] = useState<string>("");

    useEffect(() => {
        const getImageUrl = async (imageRef: StorageReference) => {
            try {
                const url = await getDownloadURL(imageRef);
                setPrevUrl(url);
            } catch (e) {
                setError(`${e}`);
            }
        };
    
            const imageRef = ref(storage, `images/${image.fileName}`);
            getImageUrl(imageRef);
        }, [image.fileName]);
    
    return (
        <div>
            <p>
                <img className="image-preview"
                    src={prevUrl}
                    alt={error}
                    
                />
            
            </p>
            <p>{image.text}</p>
            <p>{formatTimestamp(image.timestamp.toDate())}</p>
        </div>
    );
};

export default PreviewImage;