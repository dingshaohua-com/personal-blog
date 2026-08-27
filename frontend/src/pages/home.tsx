import { Link } from 'react-router';
import HelloWorld from '@/components/hello-world';

export default function Home() {
  return (
    <div className="home">
      <Link to="/about"><button>about</button></Link>
      <div>
        <div>呵呵</div>
        <HelloWorld />
      </div>
    </div>
  );
}