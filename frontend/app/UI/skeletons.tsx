// Loading animation
const shimmer =
  'before:absolute before:inset-0 before:-translate-x-full before:animate-[shimmer_2s_infinite] before:bg-gradient-to-r before:from-transparent before:via-white/60 before:to-transparent';

export function UserCardSkeleton() {
  return (
    <div className="relative overflow-hidden rounded-lg p-4 shadow-sm bg-gray-100">
      <div className="flex flex-col items-center">
        <div className={`${shimmer} w-40 h-40 rounded-full bg-gray-200 mb-4`} />
        <div className="w-full bg-gray-300 rounded-xl p-6">
          <div className="mb-1 p-4">
            <div className={`${shimmer} h-6 w-20 rounded-md bg-gray-200 mb-2`} />
            <div className={`${shimmer} h-4 w-full rounded-md bg-gray-200`} />
          </div>

          <div className="mb-1 p-4">
            <div className={`${shimmer} h-6 w-20 rounded-md bg-gray-200 mb-2`} />
            <div className={`${shimmer} h-4 w-full rounded-md bg-gray-200`} />
          </div>

          <div className="mb-1 p-4">
            <div className={`${shimmer} h-6 w-20 rounded-md bg-gray-200 mb-2`} />
            <div className={`${shimmer} h-4 w-full rounded-md bg-gray-200`} />
          </div>

          <div className="mb-1 p-4">
            <div className={`${shimmer} h-6 w-20 rounded-md bg-gray-200 mb-2`} />
            <div className={`${shimmer} h-4 w-full rounded-md bg-gray-200`} />
          </div>
        </div>
      </div>
    </div>
  );
}

