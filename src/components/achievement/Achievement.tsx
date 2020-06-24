import React from "react";
import AchievementInterface from "./AchievementInterface";

class Achievement extends React.Component<AchievementInterface>{

    render() {
        return (
            <div className="col-xs-4">
                {this.props.achievementText}
            </div>
        )
    }
}

export default Achievement