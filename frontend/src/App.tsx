/**
 * アプリ全体に共通するレイアウトや状態管理、CSSスタイルを設定する
 */

import "./App.css"
import React from "react";
import Home from "./pages/home_page/Home";

const App: React.FC = () => {
    return (
        <div>
            <Home />
        </div>
    );
};

export default App;