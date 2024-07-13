// components/UploadAvatar.tsx
import { useState } from 'react';
import { getDownloadURL, ref, uploadBytesResumable } from 'firebase/storage';
import { storage } from '../shared/firebaseconfig';
import { axiosInstance } from '../fetcher/fetcher';


function UploadAvatar({ setAvatarUrl }:
    { setAvatarUrl: (url: string) => void }) {

    const [file, setFile] = useState<File | null>(null);
    const [progress, setProgress] = useState<number>(0);


    function  handleUpload  ()  {
        if (!file) return;

        const storageRef = ref(storage, `avatars/${file.name}`);
        const uploadTask = uploadBytesResumable(storageRef, file);

        uploadTask.on(
            'state_changed',
            (snapshot) => {
                const progress = (snapshot.bytesTransferred / snapshot.totalBytes) * 100;
                setProgress(progress);
            },
            (error) => {
                console.error('Upload failed:', error);
            },
            () => {
                getDownloadURL(uploadTask.snapshot.ref).then((downloadURL) => {
                    setAvatarUrl(downloadURL);
                    // Send the URL to the backend to update the user profile
                    axiosInstance.post(`/users/profile`, {
                        "profile_image": downloadURL,
                    }, {
                        withCredentials: true,
                        headers: {
                            'Content-Type': 'application/json',
                        },
                    })
                        .then((response) => {
                            console.log('Avatar updated successfully');
                        })
                        .catch((error) => {
                            console.error('Failed to update avatar', error);
                        });
                });
            }
        );
    };

    return (
        <div>
            <input type="file" onChange={(e) => setFile(e.target.files ? e.target.files[0] : null)} />
            <button onClick={handleUpload}>Upload</button>
            <div>Progress: {progress}%</div>
        </div>
    );
};

export default UploadAvatar;
