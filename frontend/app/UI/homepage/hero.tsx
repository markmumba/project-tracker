import React from 'react';
import Image from 'next/image';
import heroImage from '/public/images/alvaro-reyes-qWwpHwip31M-unsplash.jpg';
import Link from 'next/link';

function Hero() {
    return (
        <div className="relative h-screen w-full">
            <Image
                src={heroImage}
                alt="Background Image"
                className="object-cover h-full w-full"
            />
            <div className="absolute h-screen inset-0 bg-gradient-to-r from-black  to-transparent"></div>
            <div className="absolute inset-0 flex items-center ">
                <div className="text-white p-8 max-w-4xl">
                    <h1 className="text-4xl md:text-6xl font-bold mb-4">ProjectTracker: Bridging the Gap Between Students and Lecturers</h1>
                    <p className="text-lg md:text-xl mb-8">
                        an application designed to enhance communication between lecturers
                        and students during their final project period. Students submit progress
                        documents for lecturers to review and provide feedback, while the app tracks
                        all interactions and maintains a communication history.
                    </p>
                    <Link href="/register">
                        <button className="bg-indigo-500 hover:bg-indigo-600 text-white py-2 px-4 rounded">Get Started</button>
                    </Link>
                </div>
            </div>
        </div>
    );
};

export default Hero;
