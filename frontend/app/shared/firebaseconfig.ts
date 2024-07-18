// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { getStorage } from "firebase/storage";

// TODO: Add SDKs for Firebase products that you want to use
const firebaseConfig = {
  apiKey: "AIzaSyAdylr89fILGrjg3QdnV4wp0y8dVH-XClk",
  authDomain: "projecttracker-3e7f5.firebaseapp.com",
  projectId: "projecttracker-3e7f5",
  storageBucket: "projecttracker-3e7f5.appspot.com",
  messagingSenderId: "597391387587",
  appId: "1:597391387587:web:ff53730e24dd4bc1672b76",
  measurementId: "G-9CNXMRZJ9K"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
const storage = getStorage(app);

export { storage };