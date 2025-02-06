/**
 * ダッシュボード（webサイトを開いた時に一番最初に出てくる）の画面を実現
 * ここから大まかな画面遷移を行う
 */

import React from 'react';
import { useNavigate } from "react-router-dom";
import PreviewImage from './PreviewImage';
import useImages from "../../hooks/useImages";
import './Home.css';

const Home: React.FC = () => {
    const { allImages, imagesError } = useImages();
    const navigate = useNavigate()

    const handleMyProfile = () => {
        navigate('/my-profile')
    }

    return (
        <div className="container">
            ホーム画面です。
            <div>
                <button onClick={handleMyProfile}>マイページへ</button>
            </div>
            
            <div>
                <h3>投稿</h3>
            </div>
            
            <h3>投稿一覧</h3>

            <p style={{ color: "red" }}>{imagesError && imagesError}</p>

            <div className="row">
                {allImages.map((image, index) => (
                    <div className="col-4" key={index}>
                        <PreviewImage image={image} />
                    </div>
                ))}
            </div>
        </div>
    );
};

export default Home;