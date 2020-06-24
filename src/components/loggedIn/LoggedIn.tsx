import React from "react";
import Achievement from "../achievement/Achievement";
import AchievementInterface from "../achievement/AchievementInterface";

class LoggedIn extends React.Component {
    state = {
        achievements: [],
        loading: true,
        error: false,
    }

    async componentDidMount() {
        // let headers = new Headers();
        // headers.append("Pragma", "no-cache");
        //
        // await fetch(
        //     "http://localhost:8080/api/achievements",
        //     {
        //         method: 'GET',
        //         redirect: 'follow',
        //         headers: headers
        //     })
        //     .then(res => res.json())
        //     .then(res => this.setState({
        //         achievements: res.achievements,
        //         loading: false
        //     }))
        //     .then(res => console.log())
        //     .catch((error: string) => console.log('error', error));

        fetch("http://localhost:8080/api/achievements")
            .then(res => res.json())
            .then(res => {
                    this.setState({
                        achievements: res.achievements,
                        loading: false
                    });
                },
                (error) => {
                    this.setState({
                        loading: false,
                        error
                    });
                })
    }

    render() {
        const {achievements} = this.state;

        return (
            <div className="container">
                <br />
                <span className="pull-right">
          {/*<a onClick={this.logout}>Log out</a>*/}
        </span>
                <h2>Achievements Blog</h2>
                <p>Share your achievements and comments others`</p>
                <div className="row">
                    <div className="container">
                        {achievements.map(function(achievement: AchievementInterface, i: number) {
                            return <Achievement key={i} id={achievement.id} title={achievement.title} achievementText={achievement.achievementText} authorID={achievement.authorID} date={achievement.date}/>;
                        })}
                    </div>
                </div>
            </div>
        );
    }

}


export default LoggedIn