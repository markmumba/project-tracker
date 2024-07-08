


export const truncateDescription = (description: string, maxLength: number) => {
    if (description.length <= maxLength) {
        return description;
    } else {
        return description.substring(0, maxLength) + '...';
    }
};