/**
 * マイプロフィールの画面の実現
 * ユーザ認証を得てから、マイプロフィールに遷移することができる
 */

import React from 'react';
import { useNavigate } from "react-router-dom";
import './MyProfile.css';

const MyProfile: React.FC = () => {
    const navigate = useNavigate()
    const handleHome = () => {
        navigate('/')
    }

    return (
        <div className="my-profile">
            マイページです。
            <button onClick={handleHome}>ホームへ</button>
        </div>
    );
};

export default MyProfile;