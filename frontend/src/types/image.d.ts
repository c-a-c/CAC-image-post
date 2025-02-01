import {Timestamp} from "firebase/firestore";

interface Image {
    id: string;
    fileName: string;
    text: number;
    timestamp: Timestamp;
    tags: string[];
}

export default Image;