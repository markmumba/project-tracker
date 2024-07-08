import Image from 'next/image';
import screenshot from '/public/images/screenshot.jpg';


function AboutSection() {
    return (
        <div className="flex flex-col items-center justify-center py-16 bg-gray-100 relative">
            <div className="relative max-w-7xl overflow-hidden rounded-xl">
                <Image
                    src={screenshot}
                    alt="Dashboard"
                    className="w-full h-full rounded-xl object-cover"
                />
                
            </div>
        </div>
    );
};

export default AboutSection;