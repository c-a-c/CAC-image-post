/**
 * 画面遷移のルートの設定を行う
 */

import React from "react";
import { Routes, Route } from "react-router-dom";
import Home from "../../pages/top_page/Top"
import MyProfile from "../../pages/my_page/MyProfile";

export const AppRoutes = () => {
    return (
        <Routes>
            <Route path="/" element= {<Home />} />                   {/* ホーム画面へのルート */}
            <Route path="/my-profile" element={<MyProfile />} />     {/* プロフィール画面へのルート */}
        </Routes>
    )
}